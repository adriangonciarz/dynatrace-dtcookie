package rules

type Attributes string

var Attributess = struct {
	PgName                Attributes
	PgTag                 Attributes
	ServiceManagementZone Attributes
	ServiceName           Attributes
	ServiceTag            Attributes
	ServiceType           Attributes
}{
	Attributes("PG_NAME"),
	Attributes("PG_TAG"),
	Attributes("SERVICE_MANAGEMENT_ZONE"),
	Attributes("SERVICE_NAME"),
	Attributes("SERVICE_TAG"),
	Attributes("SERVICE_TYPE"),
}

type ServiceTypes string

var ServiceTypess = struct {
	Amp                  ServiceTypes
	Cics                 ServiceTypes
	Cicsinteraction      ServiceTypes
	Customapplication    ServiceTypes
	Database             ServiceTypes
	Enterpriseservicebus ServiceTypes
	External             ServiceTypes
	Ims                  ServiceTypes
	Imsinteraction       ServiceTypes
	Messaging            ServiceTypes
	Method               ServiceTypes
	Mobile               ServiceTypes
	Process              ServiceTypes
	Queueinteraction     ServiceTypes
	Queuelistener        ServiceTypes
	Remotecall           ServiceTypes
	Rmi                  ServiceTypes
	Saasvendor           ServiceTypes
	Webrequest           ServiceTypes
	Webservice           ServiceTypes
	Website              ServiceTypes
	Zosconnect           ServiceTypes
}{
	ServiceTypes("AMP"),
	ServiceTypes("CICS"),
	ServiceTypes("CICSInteraction"),
	ServiceTypes("CustomApplication"),
	ServiceTypes("Database"),
	ServiceTypes("EnterpriseServiceBus"),
	ServiceTypes("External"),
	ServiceTypes("IMS"),
	ServiceTypes("IMSInteraction"),
	ServiceTypes("Messaging"),
	ServiceTypes("Method"),
	ServiceTypes("Mobile"),
	ServiceTypes("Process"),
	ServiceTypes("QueueInteraction"),
	ServiceTypes("QueueListener"),
	ServiceTypes("RemoteCall"),
	ServiceTypes("RMI"),
	ServiceTypes("SaasVendor"),
	ServiceTypes("WebRequest"),
	ServiceTypes("WebService"),
	ServiceTypes("WebSite"),
	ServiceTypes("ZOSConnect"),
}
