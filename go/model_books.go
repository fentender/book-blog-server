/*
 * Book Blog API
 *
 * This is a blog about books.
 *
 * API version: 0.0.7
 * Generated by: Swagger Codegen (https://github.com/swagger-api/swagger-codegen.git)
 */

package swagger

type Books struct {

	Num int32 `json:"num"`

	Books []Book `json:"books"`
}
