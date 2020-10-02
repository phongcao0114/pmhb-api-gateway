package utils

import (
	"context"
	"crypto/rand"
	"encoding/json"
	"errors"
	"fmt"
	"net/http"
	"pmhb-api-gateway/internal/kerrors"
	"pmhb-api-gateway/internal/pkg/mapper"

	"github.com/graphql-go/graphql/language/ast"

	"github.com/graphql-go/graphql"

	"time"
)

var (
	// ResponseAppID returns appID from system
	ResponseAppID string

	// BKKLocation contains time location
	BKKLocation *time.Location

	appIDKey  = "request_app_id"
	reqIDKey  = "request_id"
	dateIDKey = "request_datetime"

	// LogKey contains log timing details
	LogKey = "log_request"
)

// GetRequestID function returns request ID
func GetRequestID(ctx context.Context) string {
	requestID, _ := ctx.Value(reqIDKey).(string)
	return requestID
}

// GetAppID function returns app ID
func GetAppID(ctx context.Context) string {
	appID, _ := ctx.Value(appIDKey).(string)
	return appID
}

// GetRequestDate function returns request date
func GetRequestDate(ctx context.Context) string {
	date, _ := ctx.Value(dateIDKey).(string)
	return date
}

// ParsePagination return current page and total page
func ParsePagination(limit, total, offset int) (currentPage int, totalPage int) {
	if limit == 0 {
		limit = 20
	}

	currentPage = offset/limit + 1
	return currentPage, total
}

// DecodeToBody function is decoding all general requests from all API
func DecodeToBody(kerror *kerrors.KError, body interface{}, r *http.Request) (err error) {
	var requestInfo interface{}
	defer r.Body.Close()
	if err = json.NewDecoder(r.Body).Decode(&requestInfo); err != nil {
		err = kerror.Wrap(err, kerrors.CannotDecodeInputRequest, nil)
		return
	}
	if err = mapper.ConvertMapToModel(body, requestInfo); err != nil {
		err = kerror.Wrap(err, kerrors.MarshalFail, nil)
		return
	}
	return
}

// ConvertToModel is converting byte data to model
func ConvertToModel(data interface{}, model interface{}) error {
	byte, err := json.Marshal(data)
	if err != nil {
		return err
	}
	if err = json.Unmarshal(byte, model); err != nil {
		return err
	}
	return nil
}

func strPtr(s string) *string {
	return &s
}

func boolPtr(b bool) *bool {
	return &b
}

func GenerateTransID() string {
	b := make([]byte, 10)
	_, err := rand.Read(b)
	if err != nil {
		return ""
	}
	return fmt.Sprintf("%x", b)
}

func GetPtrIntData(input *int) int {
	if input == nil {
		return 0
	}
	return *input
}

func GetPtrStrData(input *string) string {
	if input == nil {
		return ""
	}
	return *input
}

func GetPtrFloat64Data(input *float64) float64 {
	if input == nil {
		return 0.0
	}
	return *input
}

func SplitCardDate(month, year string) string {
	if len(month) == 0 || len(year) == 0 {
		return ""
	}
	arrStr := []rune(year)
	return string(arrStr[len(arrStr)-2:]) + month
}

func HandleRespError(resp []byte) error {
	respHeader := KbankResponseHeader{}
	json.Unmarshal(resp, &respHeader)
	if respHeader != (KbankResponseHeader{}) {
		return errors.New(respHeader.Errors.ErrorDesc)
	}
	return nil
}

func ChildrenOfField(params graphql.ResolveParams) ([]string, error) {
	fields := params.Info.FieldASTs

	if len(fields) != 1 {
		return nil, errors.New(fmt.Sprintf("found more than one (%v) field ASTs at top level; unsupported behavior", len(fields)))
	}

	fields = getChildren(fields[0])

	var selected []string
	for _, field := range fields {
		selected = append(selected, field.Name.Value)
	}

	return selected, nil
}

func getChildren(field *ast.Field) []*ast.Field {
	fields := make([]*ast.Field, 0)
	selections := field.SelectionSet.Selections

	for _, selection := range selections {
		fields = append(fields, selection.(*ast.Field))
	}

	return fields
}

//func GetSelectedFields(selectionPath []string,
//	resolveParams graphql.ResolveParams) []string {
//	fields := resolveParams.Info.FieldASTs
//	for _, propName := range selectionPath {
//		found := false
//		for _, field := range fields {
//			if field.Name.Value == propName {
//				selections := field.SelectionSet.Selections
//				fields = make([]*ast.Field, 0)
//				for _, selection := range selections {
//					fields = append(fields, selection.(*ast.Field))
//				}
//				found = true
//				break
//			}
//		}
//		if !found {
//			return []string{}
//		}
//	}
//	var collect []string
//	for _, field := range fields {
//		collect = append(collect, field.Name.Value)
//	}
//	return collect
//}

func GetSelectedFields(params graphql.ResolveParams) ([]string, error) {
	fieldASTs := params.Info.FieldASTs
	if len(fieldASTs) == 0 {
		return nil, fmt.Errorf("getSelectedFields: ResolveParams has no fields")
	}
	return selectedFieldsFromSelections(params, fieldASTs[0].SelectionSet.Selections)
}

func selectedFieldsFromSelections(params graphql.ResolveParams, selections []ast.Selection) ([]string, error) {
	var selected []string
	for _, s := range selections {
		switch t := s.(type) {
		case *ast.Field:
			selected = append(selected, s.(*ast.Field).Name.Value)
		case *ast.FragmentSpread:
			n := s.(*ast.FragmentSpread).Name.Value
			frag, ok := params.Info.Fragments[n]
			if !ok {
				return nil, fmt.Errorf("getSelectedFields: no fragment found with name %v", n)
			}
			sel, err := selectedFieldsFromSelections(params, frag.GetSelectionSet().Selections)
			if err != nil {
				return nil, err
			}
			selected = append(selected, sel...)
		default:
			return nil, fmt.Errorf("getSelectedFields: found unexpected selection type %v", t)
		}
	}
	return selected, nil
}
