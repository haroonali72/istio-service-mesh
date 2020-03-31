package types

import "time"

type ServiceTemplate struct {
	Replicas          int                  `json:"replicas" bson:"replicas"`
	CompanyId         *string              `json:"company_id,omitempty" bson:"company_id" valid:"-"`
	Id                interface{}          `json:"_id,omitempty" bson:"_id" valid:"-"`
	ServiceId         *string              `json:"service_id,omitempty" bson:"service_id" binding:"required" valid:"alphanumspecial,length(6|30)~service_id must contain between 6 and 30 characters,lowercase~lowercase alphanumeric characters are allowed,required~service_id is missing in request"`
	Name              *string              `json:"name,omitempty"  bson:"name" binding:"required" valid:"alphanumspecial,length(5|30),lowercase~lowercase alphanumeric characters are allowed,required"`
	Version           *string              `json:"version,omitempty" bson:"version"  binding:"required" valid:"alphanumspecial,length(1|10),lowercase~lowercase alphanumeric characters are allowed,required"`
	ServiceType       *string              `json:"service_type,omitempty" bson:"service_type"  binding:"required" valid:"alphanumspecial,required"`
	ServiceSubType    *string              `json:"service_sub_type,omitempty" bson:"service_sub_type" valid:"-"`
	GroupId           *string              `json:"group_id,omitempty" bson:"group_id" valid:"-"`
	IsYaml            bool                 `json:"is_yaml,omitempty" bson:"is_yaml" valid:"-"`
	Deleted           bool                 `json:"deleted,omitempty" bson:"deleted" valid:"-"`
	ServiceAttributes interface{}          `json:"service_attributes,omitempty" bson:"service_attributes" binding:"required" valid:"-"`
	Status            *string              `json:"status,omitempty" bson:"status" valid:"-"`
	NumberOfInstances *int                 `json:"number_of_instances,omitempty" bson:"number_of_instances" `
	CreationDate      time.Time            `json:"creation_date,omitempty" bson:"creation_date" valid:"-"`
	NameSpace         *string              `json:"namespace,omitempty" valid:"-"`
	NameSpaceColor    *string              `json:"namespace_color,omitempty" bson:"namespace_color" valid:"-"`
	ChunckId          *string              `json:"pool_id,omitempty" valid:"-"`
	BeforeServices    []*string            `json:"in,omitempty" bson:"in" valid:"-"`
	AfterServices     []*string            `json:"out,omitempty" bson:"out" valid:"-"`
	Dependencies      map[string][]*string `json:"-" bson:"-" valid:"-"`
	Type              *string              `json:"type,omitempty" bson:"type" valid:"-"`
	Attrs             *struct {
		Body struct {
			RefWidth        string  `json:"ref_width" bson:"ref_width" valid:"-"`
			RefHeight       string  `json:"ref_height" bson:"ref_height" valid:"-"`
			Stroke          string  `json:"stroke" bson:"stroke" valid:"-"`
			Fill            string  `json:"fill" bson:"fill" valid:"-"`
			StrokeWidth     float64 `json:"stroke_width" bson:"stroke_width" valid:"-"`
			StrokeDasharray string  `json:"stroke_dasharray" bson:"stroke_dasharray" valid:"-"`
		} `json:"body,omitempty" bson:"body" valid:"-"`
		Image struct {
			RefWidth            string `json:"ref_width,omitempty" bson:"ref_width" valid:"-"`
			RefHeight           int    `json:"ref_height,omitempty" bson:"ref_height" valid:"-"`
			X                   int    `json:"x,omitempty" bson:"x" valid:"-"`
			Y                   int    `json:"y,omitempty" bson:"y" valid:"-"`
			PreserveAspectRatio string `json:"preserve_aspect_ratio,omitempty" bson:"preserve_aspect_ratio" valid:"-"`
			XlinkHref           string `json:"xlink_href,omitempty" bson:"xlink_href" valid:"-"`
		} `json:"image,omitempty" bson:"image" valid:"-"`
		Label struct {
			TextVerticalAnchor string `json:"text_vertical_anchor" bson:"text_vertical_anchor" valid:"-"`
			TextAnchor         string `json:"text_anchor" bson:"text_anchor" valid:"-"`
			RefX               string `json:"ref_x" bson:"ref_x" valid:"-"`
			RefX2              int    `json:"ref_x2" bson:"ref_x2" valid:"-"`
			RefY               int    `json:"ref_y" bson:"ref_y" valid:"-"`
			FontSize           int    `json:"font_size" bson:"font_size" valid:"-"`
			Fill               string `json:"fill" bson:"fill" valid:"-"`
			TextWrap           struct {
				Text     string `json:"text" bson:"text" valid:"-"`
				Width    int    `json:"width" bson:"width" valid:"-"`
				Height   int    `json:"height" bson:"height" valid:"-"`
				Ellipsis bool   `json:"ellipsis" bson:"ellipsis" valid:"-"`
			} `json:"text_wrap,omitempty" bson:"text_wrap" valid:"-"`
			FontFamily  string  `json:"font_family" bson:"font_family" valid:"-"`
			FontWeight  string  `json:"font_weight" bson:"font_weight" valid:"-"`
			StrokeWidth float64 `json:"stroke_width" bson:"stroke_width" valid:"-"`
		} `json:"label,omitempty" bson:"label"valid:"-"`
		ServiceType struct {
			TextVerticalAnchor string `json:"text_vertical_anchor" bson:"text_vertical_anchor" valid:"-"`
			TextAnchor         string `json:"text_anchor" bson:"text_anchor" valid:"-"`
			RefX               string `json:"ref_x" bson:"ref_x" valid:"-"`
			RefX2              int    `json:"ref_x2" bson:"ref_x2" valid:"-"`
			RefY               int    `json:"ref_y" bson:"ref_y" valid:"-"`
			FontSize           int    `json:"font_size" bson:"font_size" valid:"-"`
			Fill               string `json:"fill" bson:"fill" valid:"-"`
			TextWrap           struct {
				Text     string `json:"text" bson:"text" valid:"-"`
				Width    int    `json:"width" bson:"width" valid:"-"`
				Height   int    `json:"height" bson:"height" valid:"-"`
				Ellipsis bool   `json:"ellipsis" bson:"ellipsis" valid:"-"`
			} `json:"text_wrap,omitempty" bson:"text_wrap" valid:"-"`
			FontFamily  string  `json:"font_family" bson:"font_family" valid:"-"`
			FontWeight  string  `json:"font_weight" bson:"font_weight" valid:"-"`
			StrokeWidth float64 `json:"c" bson:"stroke_width" valid:"-"`
		} `json:"service_type,omitempty" bson:"service_type" valid:"-"`
		Root struct {
			DataTooltip                 string `json:"data_tooltip" bson:"data_tooltip" valid:"-"`
			DataTooltipPosition         string `json:"data_tooltip_position" bson:"data_tooltip_position"valid:"-"`
			DataTooltipPositionSelector string `json:"data_tooltip_position_selector" bson:"data_tooltip_position_selector" valid:"-"`
		} `json:"root,omitempty" bson:"root" valid:"-"`
	} `json:"attrs,omitempty" bson:"attrs" valid:"-"`
	Position *struct {
		X int `json:"x" bson:"x" valid:"-"`
		Y int `json:"y" bson:"y" valid:"-"`
	} `json:"position,omitempty" bson:"position" valid:"-"`
	Size *struct {
		Width  int `json:"width" bson:"width" valid:"-"`
		Height int `json:"height" bson:"height" valid:"-"`
	} `json:"size,omitempty" bson:"size" valid:"-"`
	Angle *int `json:"angle,omitempty" bson:"angle" valid:"-"`
}
