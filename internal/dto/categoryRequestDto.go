package dto


type CreateCategoryRequest struct {
	Name			string		`json:"name"`
	ParentID		uint		`json:"parent_id"`
	ImageURL		string		`json:"image_url"`
	DisplayOrder	int			`json:"display_order"`
}