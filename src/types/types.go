package types

// Request Body
type OrderRequest struct {
	Items     int   `json:"items"`
	PackSizes []int `json:"pack_sizes"`
}

// Response Body
type PacketResponse struct {
	Packs map[int]int `json:"packs"`
}
