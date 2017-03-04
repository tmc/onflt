package onflt

// WorkerSubgraph describes outstanding tasks and workers.
type WorkerSubgraph struct {
	Tasks   []Task   `json:"tasks,omitempty"`
	Workers []Worker `json:"workers,omitempty"`
}

// Task is an Onfleet Task
type Task struct {
	Attachments []*Attachment `json:"attachments,omitempty"`
	CommConfig  struct {
		DoProxy bool `json:"doProxy,omitempty"`
	} `json:"commConfig,omitempty"`
	CompleteAfter     float64 `json:"completeAfter,omitempty"`
	CompleteBefore    float64 `json:"completeBefore,omitempty"`
	CompletionDetails struct {
		Events []struct {
			Name string `json:"name,omitempty"`
			Time int64  `json:"time,omitempty"`
		} `json:"events,omitempty"`
		FailureReason     string `json:"failureReason,omitempty"`
		Notes             string `json:"notes,omitempty"`
		PhotoUploadID     string `json:"photoUploadId,omitempty"`
		SignatureUploadID string `json:"signatureUploadId,omitempty"`
		Success           bool   `json:"success,omitempty"`
		Time              int64  `json:"time,omitempty"`
	} `json:"completionDetails,omitempty"`
	Computed struct {
		Arrival struct {
			Location []float64 `json:"location,omitempty"`
			Time     string    `json:"time,omitempty"`
		} `json:"arrival,omitempty"`
		Departure struct {
			Location []float64 `json:"location,omitempty"`
			Time     string    `json:"time,omitempty"`
		} `json:"departure,omitempty"`
		Distance      float64   `json:"distance,omitempty"`
		Duration      float64   `json:"duration,omitempty"`
		FirstLocation []float64 `json:"firstLocation,omitempty"`
		LastLocation  []float64 `json:"lastLocation,omitempty"`
		Version       float64   `json:"version,omitempty"`
	} `json:"computed,omitempty"`
	Creator      string        `json:"creator,omitempty"`
	DelayTime    interface{}   `json:"delayTime,omitempty"`
	Dependencies []interface{} `json:"dependencies,omitempty"`
	Destination  struct {
		Address struct {
			Apartment  string      `json:"apartment,omitempty"`
			City       string      `json:"city,omitempty"`
			Country    string      `json:"country,omitempty"`
			Name       interface{} `json:"name,omitempty"`
			Number     string      `json:"number,omitempty"`
			PostalCode string      `json:"postalCode,omitempty"`
			State      string      `json:"state,omitempty"`
			Street     string      `json:"street,omitempty"`
		} `json:"address,omitempty"`
		CreatedByLocation bool      `json:"createdByLocation,omitempty"`
		ID                string    `json:"id,omitempty"`
		Location          []float64 `json:"location,omitempty"`
		Notes             string    `json:"notes,omitempty"`
		Organization      string    `json:"organization,omitempty"`
		TimeCreated       int64     `json:"timeCreated,omitempty"`
		TimeLastModified  int64     `json:"timeLastModified,omitempty"`
		WasGeocoded       bool      `json:"wasGeocoded,omitempty"`
	} `json:"destination,omitempty"`
	Executor       string        `json:"executor,omitempty"`
	Feedback       []interface{} `json:"feedback,omitempty"`
	ID             string        `json:"id,omitempty"`
	IsActive       bool          `json:"isActive,omitempty"`
	IsFulfilled    bool          `json:"isFulfilled,omitempty"`
	LastModifiedBy string        `json:"lastModifiedBy,omitempty"`
	Merchant       string        `json:"merchant,omitempty"`
	Notes          string        `json:"notes,omitempty"`
	Organization   string        `json:"organization,omitempty"`
	Overrides      struct {
		RecipientName                 string `json:"recipientName,omitempty"`
		RecipientNotes                string `json:"recipientNotes,omitempty"`
		RecipientSkipSMSNotifications bool   `json:"recipientSkipSMSNotifications,omitempty"`
	} `json:"overrides,omitempty"`
	PickupTask     bool    `json:"pickupTask,omitempty"`
	Quantity       float64 `json:"quantity,omitempty"`
	RecipientNotes []struct {
		ID    string `json:"id,omitempty"`
		Notes string `json:"notes,omitempty"`
	} `json:"recipientNotes,omitempty"`
	Recipients       []Recipient `json:"recipients,omitempty"`
	ServiceTime      int64       `json:"serviceTime,omitempty"`
	ShortID          string      `json:"shortId,omitempty"`
	TimeCreated      int64       `json:"timeCreated,omitempty"`
	TimeLastModified int64       `json:"timeLastModified,omitempty"`
	Worker           string      `json:"worker,omitempty"`
}

// Worker is an Onfleet Worker.
type Worker struct {
	ActiveTask          string      `json:"activeTask,omitempty"`
	Capacity            int64       `json:"capacity,omitempty"`
	CompletedTasksCount int64       `json:"completedTasksCount,omitempty"`
	DelayTime           interface{} `json:"delayTime,omitempty"`
	ID                  string      `json:"id,omitempty"`
	Image               string      `json:"image,omitempty"`
	IsResponding        bool        `json:"isResponding,omitempty"`
	Location            []float64   `json:"location,omitempty"`
	Name                string      `json:"name,omitempty"`
	OnDuty              bool        `json:"onDuty,omitempty"`
	Organization        string      `json:"organization,omitempty"`
	Password            struct {
		IsPermanent bool `json:"isPermanent,omitempty"`
	} `json:"password,omitempty"`
	Phone            string      `json:"phone,omitempty"`
	ProgressString   string      `json:"progressString,omitempty"`
	Schedule         interface{} `json:"schedule,omitempty"`
	Tasks            []string    `json:"tasks,omitempty"`
	TimeCreated      int64       `json:"timeCreated,omitempty"`
	TimeLastModified int64       `json:"timeLastModified,omitempty"`
	TimeLastSeen     int64       `json:"timeLastSeen,omitempty"`
	UserData         UserData    `json:"userData,omitempty"`
	Vehicle          struct {
		Color            string `json:"color,omitempty"`
		Description      string `json:"description,omitempty"`
		ID               string `json:"id,omitempty"`
		LicensePlate     string `json:"licensePlate,omitempty"`
		TimeLastModified int64  `json:"timeLastModified,omitempty"`
		Type             string `json:"type,omitempty"`
	} `json:"vehicle,omitempty"`
}

// Attachment is an file attached to a Task.
type Attachment struct {
	AttachmentID   string `json:"attachmentId,omitempty"`
	AttachmentType string `json:"attachmentType,omitempty"`
	Status         string `json:"status,omitempty"`
}

// Recipient is a delivery recipient.
type Recipient struct {
	ID                   string  `json:"id,omitempty"`
	Name                 string  `json:"name,omitempty"`
	Organization         string  `json:"organization,omitempty"`
	Phone                string  `json:"phone,omitempty"`
	SkipSMSNotifications bool    `json:"skipSMSNotifications,omitempty"`
	TimeCreated          float64 `json:"timeCreated,omitempty"`
	TimeLastModified     float64 `json:"timeLastModified,omitempty"`
}

// UserData contains data about a worker's device.
type UserData struct {
	AppVersion        string  `json:"appVersion,omitempty"`
	BatteryLevel      float64 `json:"batteryLevel,omitempty"`
	DeviceDescription string  `json:"deviceDescription,omitempty"`
	Platform          string  `json:"platform,omitempty"`
}
