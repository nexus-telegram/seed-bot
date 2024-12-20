package types

import "time"

type BirdsHappiness struct {
	Data struct {
		Id               string      `json:"id"`
		Type             string      `json:"type"`
		Status           string      `json:"status"`
		EnergyLevel      int         `json:"energy_level"`
		EnergyMax        int         `json:"energy_max"`
		HappinessLevel   int         `json:"happiness_level"`
		TaskLevel        int         `json:"task_level"`
		IsLeader         bool        `json:"is_leader"`
		HuntStartAt      time.Time   `json:"hunt_start_at"`
		HuntEndAt        time.Time   `json:"hunt_end_at"`
		ReceivedRewardAt time.Time   `json:"received_reward_at"`
		OnMarket         bool        `json:"on_market"`
		OwnerId          string      `json:"owner_id"`
		CreatedAt        time.Time   `json:"created_at"`
		UpdatedAt        time.Time   `json:"updated_at"`
		MarketId         interface{} `json:"market_id"`
		Price            interface{} `json:"price"`
	} `json:"data"`
}
type Bird struct {
	Id               string      `json:"id"`
	Type             string      `json:"type"`
	Status           string      `json:"status"`
	EnergyLevel      int         `json:"energy_level"`
	EnergyMax        int         `json:"energy_max"`
	HappinessLevel   int         `json:"happiness_level"`
	TaskLevel        int         `json:"task_level"`
	IsLeader         bool        `json:"is_leader"`
	HuntStartAt      time.Time   `json:"hunt_start_at"`
	HuntEndAt        time.Time   `json:"hunt_end_at"`
	ReceivedRewardAt time.Time   `json:"received_reward_at"`
	OnMarket         bool        `json:"on_market"`
	OwnerId          string      `json:"owner_id"`
	CreatedAt        time.Time   `json:"created_at"`
	UpdatedAt        time.Time   `json:"updated_at"`
	MarketId         interface{} `json:"market_id"`
	Price            interface{} `json:"price"`
}
type BirdData struct {
	Bird Bird `json:"data"`
}
type BirdsData struct {
	Bird []Bird `json:"data"`
}

type CompleteBirdHunterResponse struct {
	Data struct {
		Bird struct {
			Id               string      `json:"id"`
			Type             string      `json:"type"`
			Status           string      `json:"status"`
			EnergyLevel      int         `json:"energy_level"`
			EnergyMax        int         `json:"energy_max"`
			HappinessLevel   int         `json:"happiness_level"`
			TaskLevel        int         `json:"task_level"`
			IsLeader         bool        `json:"is_leader"`
			HuntStartAt      time.Time   `json:"hunt_start_at"`
			HuntEndAt        time.Time   `json:"hunt_end_at"`
			ReceivedRewardAt time.Time   `json:"received_reward_at"`
			OnMarket         bool        `json:"on_market"`
			OwnerId          string      `json:"owner_id"`
			CreatedAt        time.Time   `json:"created_at"`
			UpdatedAt        time.Time   `json:"updated_at"`
			MarketId         interface{} `json:"market_id"`
			Price            interface{} `json:"price"`
		} `json:"bird"`
		EggId        interface{} `json:"egg_id"`
		EggType      interface{} `json:"egg_type"`
		WormId       interface{} `json:"worm_id"`
		WormType     interface{} `json:"worm_type"`
		EggPieceId   interface{} `json:"egg_piece_id"`
		EggPieceType interface{} `json:"egg_piece_type"`
		SeedAmount   int         `json:"seed_amount"`
	} `json:"data"`
}
type CompleteBirdHuntPayload struct {
	BirdId string `json:"bird_id"`
}

type GetMyBirds struct {
	Data struct {
		Total    int         `json:"total"`
		Page     int         `json:"page"`
		PageSize int         `json:"page_size"`
		Birds    []BirdsData `json:"items"`
	} `json:"data"`
}
