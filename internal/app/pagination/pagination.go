package pagination

import (
	"encoding/json"
	tgbotapi "github.com/go-telegram-bot-api/telegram-bot-api/v5"
	"go-tg-bot/internal/app/path"
	"go-tg-bot/internal/service/product"
	"math"
)

const LIMIT = 5

type CallbackListData struct {
	Page int `json:"page"`
}

type Item interface {
	String() string
}

type EntityService interface {
	List(offset, limit uint) ([]product.Product, error)
	Count() int
}

type Pagination struct {
	service EntityService
}

func NewPagination(service EntityService) *Pagination {
	return &Pagination{
		service,
	}
}

func (p *Pagination) BuildListMessage(chatID int64, data CallbackListData) (*tgbotapi.MessageConfig, error) {
	page := data.Page
	offset := page * LIMIT
	maxPage := int(math.Ceil(float64(p.service.Count()) / float64(LIMIT)))

	items, err := p.service.List(uint(offset), LIMIT)
	if err != nil {
		return nil, err
	}

	var buttons []tgbotapi.InlineKeyboardButton

	if page > 0 {
		prevBtn, err := makeButton("Previous page", page-1)
		if err != nil {
			return nil, err
		}

		buttons = append(buttons, *prevBtn)
	}

	if page < maxPage-1 {
		nextBtn, err := makeButton("Next page", page+1)
		if err != nil {
			return nil, err
		}

		buttons = append(buttons, *nextBtn)
	}

	msg := tgbotapi.NewMessage(chatID, formatListItems(items))
	if len(buttons) > 0 {
		msg.ReplyMarkup = tgbotapi.NewInlineKeyboardMarkup(buttons)
	}

	return &msg, nil
}

func makeButton(text string, page int) (*tgbotapi.InlineKeyboardButton, error) {
	serializedData, err := json.Marshal(CallbackListData{
		Page: page,
	})
	if err != nil {
		return nil, err
	}

	callbackPath := path.CallbackPath{
		CallbackName: "list",
		CallbackData: string(serializedData),
	}

	newButton := tgbotapi.NewInlineKeyboardButtonData(text, callbackPath.String())

	return &newButton, nil
}

func formatListItems(items []product.Product) string {
	str := ""
	for _, item := range items {
		str += item.String() + "\n"
	}

	return str
}
