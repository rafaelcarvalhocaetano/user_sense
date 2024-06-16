package model

// Components contém informações sobre os componentes do Webhook.
type WebhookData struct {
	Object string   `json:"object,omitempty"` // Opcional. Todos os eventos do Webhook para a API do WhatsApp Cloud pertencem ao objeto whatsapp_business_account.
	Entry  []*Entry `json:"entry,omitempty"`  // Opcional. Uma matriz de objetos de entrada.
}

// EntryObject contém informações sobre o Webhook recebido.
type Entry struct {
	ID      string    `json:"id,omitempty"`      // Opcional. O ID das Contas do WhatsApp Business a que este Webhook pertence.
	Changes []*Change `json:"changes,omitempty"` // Opcional. Alterações que desencadearam a chamada do Webhook. Este campo contém uma matriz de objetos de alteração.
}

// ChangeObject representa informações sobre uma mudança relacionada a um campo específico.
type Change struct {
	Value *Value `json:"value,omitempty"` // Opcional. Um objeto de valor. Contém detalhes das alterações relacionadas ao campo especificado.
	Field string `json:"field"`           // O tipo de notificação que você está recebendo nesse Webhook. Atualmente, a única opção para esta API é "messages".
}

// ValueObject contém informações associadas ao webhook recebido.
type Value struct {
	MessagingProduct *string    `json:"messaging_product,omitempty"`
	Metadata         *Metadata  `json:"metadata,omitempty"`
	Messages         []*Message `json:"messages,omitempty"`
	Statuses         []*Status  `json:"statuses,omitempty"`
	Errors           []*Error   `json:"errors,omitempty"`
	Contacts         []Contact  `json:"contacts"`
}

type Status struct {
	ID           string        `json:"id"`
	Status       string        `json:"status"`
	Timestamp    string        `json:"timestamp"`
	RecipientID  *string       `json:"recipient_id"`
	Type         string        `json:"type"`
	Conversation *Conversation `json:"conversation,omitempty"`
	Pricing      *Pricing      `json:"pricing,omitempty"`
}

type Channel struct {
	PhoneNumber string `json:"id"`
	Status      bool   `json:"status"`
}

type Metadata struct {
	DisplayPhoneNumber   *string   `json:"display_phone_number,omitempty"`
	PhoneNumberID        *string   `json:"phone_number_id,omitempty,omitempty"`
	StickerPackID        *string   `json:"sticker-pack-id,omitempty"`
	StickerPackName      *string   `json:"sticker-pack-name,omitempty"`
	StickerPackPublisher *string   `json:"sticker-pack-publisher,omitempty"`
	Emojis               *[]string `json:"emojis,omitempty"`
	IosAppStoreLink      *string   `json:"ios-app-store-link,omitempty"`
	AndroidAppStoreLink  *string   `json:"android-app-store-link,omitempty"`
	IsFirstPartySticker  *bool     `json:"is-first-party-sticker,omitempty"`
}

// Profile representa um objeto de perfil.
type Profile struct {
	Name string `json:"name,omitempty"`
}

type MessageDispatch struct {
	MessagingProduct string `json:"messaging_product"`
	RecipientType    string `json:"recipient_type"`
	To               string `json:"to"`
	Type             string `json:"type"`
	Interactive      any    `json:"interactive"`
	Text             *Text  `json:"text"`
}

type Webhook struct {
	Messages []Message `json:"messages,omitempty"`
	Statuses []Status  `json:"statuses,omitempty"`
	Errors   []Error   `json:"errors,omitempty"`
}

// Interactive representa uma mensagem interativa.
type Interactive struct {
	Type        string  `json:"type"`
	Header      *Header `json:"header,omitempty"`
	Body        *Body   `json:"body,omitempty"`
	Footer      *Footer `json:"footer,omitempty"`
	Action      *Action `json:"action"`
	ButtonReply *Reply  `json:"button_reply,omitempty"`
	ListReply   *Reply  `json:"list_reply,omitempty"`
}

// Header representa o cabeçalho de uma mensagem.
type Header struct {
	Type     string `json:"type"`
	Text     string `json:"text,omitempty"`
	Video    *Media `json:"video,omitempty"`
	Image    *Media `json:"image,omitempty"`
	Document *Media `json:"document,omitempty"`
}

// Body representa o corpo de uma mensagem.
type Body struct {
	Text string `json:"text"`
}

// Footer representa o rodapé de uma mensagem.
type Footer struct {
	Text string `json:"text"`
}

// Action representa uma ação em uma mensagem.
type Action struct {
	Button            string     `json:"button,omitempty"`
	Buttons           []*Button  `json:"buttons,omitempty"`
	Sections          []*Section `json:"sections,omitempty"`
	CatalogID         string     `json:"catalog_id,omitempty"`
	ProductRetailerID string     `json:"product_retailer_id,omitempty"`
}

// Button representa um botão em uma mensagem.
type Button struct {
	Type    *string `json:"type,omitempty"`
	Title   string  `json:"title"`
	ID      *string `json:"id,omitempty"`
	Payload *string `json:"payload,omitempty"`
	Text    *string `json:"text,omitempty"`
}

// Section representa uma seção em uma mensagem.
type Section struct {
	Title        string     `json:"title,omitempty"`
	Rows         []*Row     `json:"rows,omitempty"`
	ProductItems []*Product `json:"product_items,omitempty"`
}

// Row representa uma linha em uma seção de lista.
type Row struct {
	ID          string `json:"id"`
	Title       string `json:"title"`
	Description string `json:"description,omitempty"`
}

// Product representa um produto em uma mensagem de vários produtos.
type Product struct {
	ProductRetailerID string  `json:"product_retailer_id"`
	Quantity          int     `json:"quantity"`
	ItemPrice         float64 `json:"item_price"`
	Currency          string  `json:"currency"`
}

// Media representa um objeto de mídia.
type Media struct {
	ID       string    `json:"id"`
	Caption  *string   `json:"caption,omitempty"`
	Filename *string   `json:"filename,omitempty"`
	MimeType string    `json:"mime_type"`
	SHA256   string    `json:"sha256"`
	File     *string   `json:"file,omitempty"`
	Metadata *Metadata `json:"metadata"`
}

// Referral representa um objeto de referência.
type Referral struct {
	SourceURL    string  `json:"source_url"`
	SourceType   string  `json:"source_type"`
	SourceID     string  `json:"source_id"`
	Headline     string  `json:"headline"`
	Body         string  `json:"body"`
	MediaType    string  `json:"media_type"`
	ImageURL     *string `json:"image_url,omitempty"`
	VideoURL     *string `json:"video_url,omitempty"`
	ThumbnailURL *string `json:"thumbnail_url,omitempty"`
	Image        *Media  `json:"image,omitempty"`
	Video        *Media  `json:"video,omitempty"`
}

// Message representa um array de objetos de mensagem.
type Message struct {
	From           string       `json:"from"`
	ID             string       `json:"id"`
	Timestamp      string       `json:"timestamp"`
	Type           string       `json:"type"`
	Context        *Context     `json:"context,omitempty"`
	Identity       *Identity    `json:"identity,omitempty"`
	Text           *Text        `json:"text,omitempty"`
	Audio          *Media       `json:"audio,omitempty"`
	Image          *Media       `json:"image,omitempty"`
	Sticker        *Media       `json:"sticker,omitempty"`
	Video          *Media       `json:"video,omitempty"`
	Interactive    *Interactive `json:"interactive,omitempty"`
	Order          *Order       `json:"order,omitempty"`
	Document       *Media       `json:"document,omitempty"`
	Errors         *[]Error     `json:"errors,omitempty"`
	System         *System      `json:"system,omitempty"`
	Button         *Button      `json:"button,omitempty"`
	Referral       *Referral    `json:"referral,omitempty"`
	MessageID      *string      `json:"message_id,omitempty"`
	SenderID       *string      `json:"sender_id,omitempty"`
	Content        *string      `json:"content,omitempty"`
	MessageObjects *[]any       `json:"message_objects,omitempty"`
	Location       *Location    `json:"location,omitempty"`
	Contacts       *[]Contact   `json:"contacts,omitempty"`
	Reaction       *Reaction    `json:"reaction,omitempty"`
}

// Location representa os detalhes de uma localização estática.
type Location struct {
	Latitude  *float64 `json:"latitude,omitempty"`
	Longitude *float64 `json:"longitude,omitempty"`
	Address   *string  `json:"address,omitempty"`
	Name      *string  `json:"name,omitempty"`
	URL       *string  `json:"url,omitempty"`
}

// Address representa um endereço de contato.
type Address struct {
	Street      *string `json:"street,omitempty"`
	City        *string `json:"city,omitempty"`
	State       *string `json:"state,omitempty"`
	Zip         *string `json:"zip,omitempty"`
	Country     *string `json:"country,omitempty"`
	CountryCode *string `json:"country_code,omitempty"`
	Type        *string `json:"type,omitempty"`
}

type Contact struct {
	Addresses []*Address    `json:"addresses,omitempty"`
	Birthday  *string       `json:"birthday,omitempty"`
	Emails    []*Email      `json:"emails,omitempty"`
	IMS       []*IMS        `json:"ims,omitempty"`
	Name      *Name         `json:"name,omitempty"`
	Org       *Organization `json:"org,omitempty"`
	Phones    []*Phone      `json:"phones,omitempty"`
	URLs      []*URL        `json:"urls,omitempty"`
	Profile   *Profile      `json:"profile,omitempty"`
	WaID      *string       `json:"wa_id,omitempty"`
}

// Name representa o nome completo do contato.
type Name struct {
	FirstName  *string `json:"first_name,omitempty"`
	MiddleName *string `json:"middle_name,omitempty"`
	LastName   *string `json:"last_name,omitempty"`
	Formatted  *string `json:"formatted_name,omitempty"`
	Prefix     *string `json:"name_prefix,omitempty"`
	Suffix     *string `json:"name_suffix,omitempty"`
}

// Email representa um endereço de e-mail de contato.
type Email struct {
	Email *string `json:"email,omitempty"`
	Type  *string `json:"type,omitempty"`
}

type IMS struct {
	Service *string `json:"service,omitempty"`
	UserID  *string `json:"user_id,omitempty"`
}

// Organization representa informações da organização de contato.
type Organization struct {
	Company    *string `json:"applications,omitempty"`
	Department *string `json:"department,omitempty"`
	Title      *string `json:"title,omitempty"`
}

// Phone representa um número de telefone de contato.
type Phone struct {
	Number *string `json:"phone,omitempty"`
	WaID   *string `json:"wa_id,omitempty"`
	Type   *string `json:"type,omitempty"`
}

// URL representa uma URL de contato.
type URL struct {
	URL  *string `json:"url,omitempty"`
	Type *string `json:"type,omitempty"`
}

// Order representa um objeto de pedido.
type Order struct {
	CatalogID    *string    `json:"catalog_id,omitempty"`
	Text         *string    `json:"text,omitempty"`
	ProductItems *[]Product `json:"product_items,omitempty"`
}

// Text representa um objeto de mensagem de texto.
type Text struct {
	Body string `json:"body"`
}

// ReferredProduct representa um objeto de produto referido.
type ReferredProduct struct {
	CatalogID         *string `json:"catalog_id,omitempty"`
	ProductRetailerID *string `json:"product_retailer_id,omitempty"`
}

// Context representa um objeto de contexto.
type Context struct {
	Forwarded           *bool            `json:"forwarded,omitempty"`
	FrequentlyForwarded *bool            `json:"frequently_forwarded,omitempty"`
	From                *string          `json:"from,omitempty"`
	ID                  *string          `json:"id,omitempty"`
	ReferredProduct     *ReferredProduct `json:"referred_product,omitempty"`
	SentMessageID       *string          `json:"sent_message_id,omitempty"`
}

// Identity representa um objeto de identidade.
type Identity struct {
	Acknowledged     *string `json:"acknowledged,omitempty"`
	CreatedTimestamp *int    `json:"created_timestamp,omitempty"`
	Hash             *string `json:"hash,omitempty"`
	UserID           *string `json:"user,omitempty"`
}

// Reaction representa um objeto de reação a uma mensagem.
type Reaction struct {
	MessageID string  `json:"message_id"`
	Emoji     *string `json:"emoji,omitempty"`
}

type Reply struct {
	ID          string  `json:"id"`
	Title       string  `json:"title"`
	Description *string `json:"description,omitempty"`
}

// SystemMessageObject é adicionado aos Webhooks se um usuário mudou seu número de telefone ou se a identidade de um usuário mudou potencialmente no WhatsApp.
type System struct {
	Body     string  `json:"body"`
	NewWaID  *string `json:"new_wa_id,omitempty"`
	Identity *string `json:"identity,omitempty"`
	Type     string  `json:"type"`
	User     *string `json:"user,omitempty"`
}

// Payment rastreia os atributos das alterações de transação iniciadas pelo usuário.
type Payment struct {
	ID        string          `json:"id"`
	From      string          `json:"from"`
	Type      string          `json:"type"`
	Status    string          `json:"status"`
	Payment   *PaymentDetails `json:"payment,omitempty"`
	Timestamp string          `json:"timestamp"`
}

// PaymentDetails contém informações detalhadas sobre o pagamento.
type PaymentDetails struct {
	ReferenceID string `json:"reference_id"`
}
