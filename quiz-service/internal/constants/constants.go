package constants

// ServiceName the name of this module/service
const ServiceName = "quizs"

// GRPC Service Names
const (
	OrganizationServiceName = "ORGANIZATION"
	QuizsServiceName        = "QUIZS"
)

// Dependency Injection Keys
const (
	RegistryKey                 = "registry"
	DomainDispatcherKey         = "domainDispatcher"
	DatabaseTransactionKey      = "tx"
	MessagePublisherKey         = "messagePublisher"
	MessageSubscriberKey        = "messageSubscriber"
	EventPublisherKey           = "eventPublisher"
	CommandPublisherKey         = "commandPublisher"
	ReplyPublisherKey           = "replyPublisher"
	SagaStoreKey                = "sagaStore"
	InboxStoreKey               = "inboxStore"
	ApplicationKey              = "app"
	DomainEventHandlersKey      = "domainEventHandlers"
	IntegrationEventHandlersKey = "integrationEventHandlers"
	CommandHandlersKey          = "commandHandlers"
	ReplyHandlersKey            = "replyHandlers"

	QuizsRepoKey = "quizsRepo"
)

// Repository Table Names
const (
	OutboxTableNames   = ServiceName + ".outbox"
	InboxTableName     = ServiceName + ".inbox"
	EventsTableName    = ServiceName + ".events"
	SnapshotsTableName = ServiceName + ".snapshots"
	SagasTableName     = ServiceName + ".sagas"
	QuizsTableName     = ServiceName + ".quizs"
)

// Metric Names
const (
	QuizsRegisteredCount = "quizs_registered_count"
)
