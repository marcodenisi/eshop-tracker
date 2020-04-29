package model

import (
	"encoding/json"
	"net/http"
	"time"
)

// EuGamesResponse represents the response of the Nintendo EU web service
type EuGamesResponse struct {
	ResponseHeader struct {
		Status int `json:"status"`
		QTime  int `json:"QTime"`
		Params struct {
			Q     string `json:"q"`
			Start string `json:"start"`
			Fq    string `json:"fq"`
			Sort  string `json:"sort"`
			Rows  string `json:"rows"`
			Wt    string `json:"wt"`
		} `json:"params"`
	} `json:"responseHeader"`
	Response struct {
		NumFound int      `json:"numFound"`
		Start    int      `json:"start"`
		Games    []EuGame `json:"docs"`
	} `json:"response"`
}

// DecodeEuGamesResponse transforms an http response body into a EuGamesResponse object
func DecodeEuGamesResponse(r *http.Response) EuGamesResponse {
	var dto EuGamesResponse
	d := json.NewDecoder(r.Body)
	d.Decode(&dto)

	return dto
}

// EuGame represents the single game retrieved from Nintendo EU web service
type EuGame struct {
	FsID                             string      `json:"fs_id"`
	ChangeDate                       time.Time   `json:"change_date"`
	URL                              string      `json:"url"`
	Type                             string      `json:"type"`
	DatesReleasedDts                 []time.Time `json:"dates_released_dts"`
	ClubNintendo                     bool        `json:"club_nintendo"`
	PrettyDateS                      string      `json:"pretty_date_s"`
	PlayModeTvModeB                  bool        `json:"play_mode_tv_mode_b"`
	PlayModeHandheldModeB            bool        `json:"play_mode_handheld_mode_b,omitempty"`
	ProductCodeTxt                   []string    `json:"product_code_txt"`
	ImageURLSqS                      string      `json:"image_url_sq_s"`
	DeprioritiseB                    bool        `json:"deprioritise_b"`
	PgS                              string      `json:"pg_s"`
	GiftFinderDetailPageImageURLS    string      `json:"gift_finder_detail_page_image_url_s,omitempty"`
	CompatibleController             []string    `json:"compatible_controller,omitempty"`
	ImageURL                         string      `json:"image_url,omitempty"`
	OriginallyForT                   string      `json:"originally_for_t"`
	PaidSubscriptionRequiredB        bool        `json:"paid_subscription_required_b,omitempty"`
	CloudSavesB                      bool        `json:"cloud_saves_b"`
	DigitalVersionB                  bool        `json:"digital_version_b"`
	TitleExtrasTxt                   []string    `json:"title_extras_txt"`
	ImageURLH2X1S                    string      `json:"image_url_h2x1_s"`
	SystemType                       []string    `json:"system_type"`
	AgeRatingSortingI                int         `json:"age_rating_sorting_i"`
	GameCategoriesTxt                []string    `json:"game_categories_txt"`
	PlayModeTabletopModeB            bool        `json:"play_mode_tabletop_mode_b,omitempty"`
	Publisher                        string      `json:"publisher"`
	ProductCodeSs                    []string    `json:"product_code_ss"`
	Excerpt                          string      `json:"excerpt"`
	NsuidTxt                         []string    `json:"nsuid_txt"`
	DateFrom                         time.Time   `json:"date_from"`
	LanguageAvailability             []string    `json:"language_availability"`
	PriceHasDiscountB                bool        `json:"price_has_discount_b,omitempty"`
	PriceDiscountPercentageF         float64     `json:"price_discount_percentage_f"`
	Title                            string      `json:"title"`
	SortingTitle                     string      `json:"sorting_title"`
	GiftFinderCarouselImageURLS      string      `json:"gift_finder_carousel_image_url_s,omitempty"`
	WishlistEmailSquareImageURLS     string      `json:"wishlist_email_square_image_url_s"`
	PlayersTo                        int         `json:"players_to"`
	WishlistEmailBanner640WImageURLS string      `json:"wishlist_email_banner640w_image_url_s"`
	VoiceChatB                       bool        `json:"voice_chat_b,omitempty"`
	PlayableOnTxt                    []string    `json:"playable_on_txt"`
	HitsI                            int         `json:"hits_i"`
	PrettyGameCategoriesTxt          []string    `json:"pretty_game_categories_txt"`
	GiftFinderWishlistImageURLS      string      `json:"gift_finder_wishlist_image_url_s,omitempty"`
	SwitchGameVoucherB               bool        `json:"switch_game_voucher_b"`
	GameCategory                     []string    `json:"game_category"`
	SystemNamesTxt                   []string    `json:"system_names_txt"`
	PrettyAgeratingS                 string      `json:"pretty_agerating_s"`
	PriceRegularF                    float64     `json:"price_regular_f,omitempty"`
	EshopRemovedB                    bool        `json:"eshop_removed_b"`
	PlayersFrom                      int         `json:"players_from,omitempty"`
	AgeRatingType                    string      `json:"age_rating_type"`
	PriceSortingF                    float64     `json:"price_sorting_f"`
	PriceLowestF                     float64     `json:"price_lowest_f"`
	AgeRatingValue                   string      `json:"age_rating_value"`
	PhysicalVersionB                 bool        `json:"physical_version_b"`
	WishlistEmailBanner460WImageURLS string      `json:"wishlist_email_banner460w_image_url_s"`
	Version                          int64       `json:"_version_"`
	Popularity                       int         `json:"popularity"`
	Internet                         bool        `json:"internet,omitempty"`
	DemoAvailability                 bool        `json:"demo_availability,omitempty"`
	PriceDiscountedF                 float64     `json:"price_discounted_f,omitempty"`
	HdRumbleB                        bool        `json:"hd_rumble_b,omitempty"`
	MultiplayerMode                  string      `json:"multiplayer_mode,omitempty"`
	IrMotionCameraB                  bool        `json:"ir_motion_camera_b,omitempty"`
	GiftFinderDescriptionS           string      `json:"gift_finder_description_s,omitempty"`
	Priority                         time.Time   `json:"priority,omitempty"`
	Developer                        string      `json:"developer,omitempty"`
	AddOnContentB                    bool        `json:"add_on_content_b,omitempty"`
	DlcShownB                        bool        `json:"dlc_shown_b,omitempty"`
}
