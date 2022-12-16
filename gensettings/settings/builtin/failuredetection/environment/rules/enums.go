package rules

type Attributes string

var Attributess = struct {
	ServiceTag            Attributes
	PgName                Attributes
	PgTag                 Attributes
	ServiceManagementZone Attributes
	ServiceName           Attributes
	ServiceType           Attributes
}{
	Attributes("SERVICE_TAG"),
	Attributes("PG_NAME"),
	Attributes("PG_TAG"),
	Attributes("SERVICE_MANAGEMENT_ZONE"),
	Attributes("SERVICE_NAME"),
	Attributes("SERVICE_TYPE"),
}

type ServiceTypes string

var ServiceTypess = struct {
	Saasvendor           ServiceTypes
	Method               ServiceTypes
	Process              ServiceTypes
	External             ServiceTypes
	Messaging            ServiceTypes
	Webrequest           ServiceTypes
	Imsinteraction       ServiceTypes
	Cics                 ServiceTypes
	Mobile               ServiceTypes
	Enterpriseservicebus ServiceTypes
	Queuelistener        ServiceTypes
	Website              ServiceTypes
	Remotecall           ServiceTypes
	Database             ServiceTypes
	Ims                  ServiceTypes
	Cicsinteraction      ServiceTypes
	Zosconnect           ServiceTypes
	Rmi                  ServiceTypes
	Webservice           ServiceTypes
	Amp                  ServiceTypes
	Queueinteraction     ServiceTypes
	Customapplication    ServiceTypes
}{
	ServiceTypes("SaasVendor"),
	ServiceTypes("Method"),
	ServiceTypes("Process"),
	ServiceTypes("External"),
	ServiceTypes("Messaging"),
	ServiceTypes("WebRequest"),
	ServiceTypes("IMSInteraction"),
	ServiceTypes("CICS"),
	ServiceTypes("Mobile"),
	ServiceTypes("EnterpriseServiceBus"),
	ServiceTypes("QueueListener"),
	ServiceTypes("WebSite"),
	ServiceTypes("RemoteCall"),
	ServiceTypes("Database"),
	ServiceTypes("IMS"),
	ServiceTypes("CICSInteraction"),
	ServiceTypes("ZOSConnect"),
	ServiceTypes("RMI"),
	ServiceTypes("WebService"),
	ServiceTypes("AMP"),
	ServiceTypes("QueueInteraction"),
	ServiceTypes("CustomApplication"),
}
