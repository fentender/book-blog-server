/*
 * Book Blog API
 *
 * This is a blog about books.
 *
 * API version: 0.0.7
 * Generated by: Swagger Codegen (https://github.com/swagger-api/swagger-codegen.git)
 */

package swagger

type Reviews struct {

	Num int32 `json:"num"`

	Reviews []Review `json:"reviews"`
}
