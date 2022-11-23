package dashboards_test

import (
	"encoding/json"
	"fmt"
	"os"
	"testing"

	"github.com/dtcookie/dynatrace/api/config/dashboards"
)

func TestValidate(t *testing.T) {
	envURL := os.Getenv("DYNATRACE_ENV_URL")
	apiToken := os.Getenv("DYNATRACE_API_TOKEN")
	// rest.Verbose = true
	client := dashboards.NewService(envURL+"/api/config/v1", apiToken)
	var dashboard dashboards.Dashboard
	if err := json.Unmarshal([]byte(dbJSON), &dashboard); err != nil {
		t.Error(err)
		return
	}
	errors := client.Validate(&dashboard)
	for _, err := range errors {
		fmt.Println(err)
	}
}

const dbJSON = `
{
    "dashboardMetadata": {
        "dynamicFilters": {
            "filters": [
                "RELATED_NAMESPACE",
                "KUBERNETES_CLUSTER"
            ]
        },
        "name": "000-1-000",
        "owner": "Dynatrace",
        "preset": true,
        "shared": true,
        "tags": [
            "Kubernetes"
        ]
    },
    "tiles": [
        {
            "assignedEntities": [],
            "bounds": {
                "height": 38,
                "left": 0,
                "top": 380,
                "width": 836
            },
            "chartVisible": false,
            "configured": true,
            "excludeMaintenanceWindows": false,
            "markdown": "## Memory requests quotas",
            "name": "Markdown",
            "tileType": "MARKDOWN"
        },
        {
            "assignedEntities": [],
            "bounds": {
                "height": 38,
                "left": 0,
                "top": 38,
                "width": 836
            },
            "chartVisible": false,
            "configured": true,
            "excludeMaintenanceWindows": false,
            "markdown": "## CPU requests quotas",
            "name": "Markdown",
            "tileType": "MARKDOWN"
        },
        {
            "assignedEntities": [],
            "bounds": {
                "height": 38,
                "left": 836,
                "top": 380,
                "width": 836
            },
            "chartVisible": false,
            "configured": true,
            "excludeMaintenanceWindows": false,
            "markdown": "## Memory limits quotas",
            "name": "Markdown",
            "tileType": "MARKDOWN"
        },
        {
            "assignedEntities": [],
            "bounds": {
                "height": 38,
                "left": 836,
                "top": 38,
                "width": 836
            },
            "chartVisible": false,
            "configured": true,
            "excludeMaintenanceWindows": false,
            "markdown": "## CPU limits quotas",
            "name": "Markdown",
            "tileType": "MARKDOWN"
        },
        {
            "assignedEntities": [],
            "bounds": {
                "height": 38,
                "left": 0,
                "top": 0,
                "width": 950
            },
            "chartVisible": false,
            "configured": true,
            "excludeMaintenanceWindows": false,
            "markdown": "This dashboard provides an overview of your [resource quotas](https://kubernetes.io/docs/concepts/policy/resource-quotas/) by namespace. If the dashboard is empty, you might not have any resource quotas set.",
            "name": "Markdown",
            "tileType": "MARKDOWN"
        },
        {
            "assignedEntities": [],
            "bounds": {
                "height": 38,
                "left": 0,
                "top": 722,
                "width": 836
            },
            "chartVisible": false,
            "configured": true,
            "excludeMaintenanceWindows": false,
            "markdown": "## Pod count quotas",
            "name": "Markdown",
            "tileType": "MARKDOWN"
        },
        {
            "assignedEntities": [],
            "bounds": {
                "height": 38,
                "left": 1368,
                "top": 38,
                "width": 304
            },
            "chartVisible": false,
            "configured": true,
            "excludeMaintenanceWindows": false,
            "markdown": "## [📝We'd love your feedback!](https://dt-url.net/k8snrqd)",
            "name": "Markdown",
            "tileType": "MARKDOWN"
        },
        {
            "assignedEntities": [],
            "bounds": {
                "height": 304,
                "left": 0,
                "top": 76,
                "width": 304
            },
            "chartVisible": false,
            "configured": true,
            "customName": "CPU requests quota used",
            "excludeMaintenanceWindows": false,
            "metricExpressions": [
                "resolution=null\u0026(builtin:kubernetes.resourcequota.requests_cpu_used:last:splitBy(\"dt.entity.cloud_application_namespace\",\"k8s.resourcequota.name\"):sum:sort(value(sum,descending))):limit(100):names:fold(auto)"
            ],
            "name": "Quota used",
            "queries": [
                {
                    "enabled": true,
                    "id": "B",
                    "metricSelector": "builtin:kubernetes.resourcequota.requests_cpu_used:last:splitBy(\"dt.entity.cloud_application_namespace\",\"k8s.resourcequota.name\"):sum:sort(value(sum,descending))",
                    "splitBy": [
                        "dt.entity.cloud_application_namespace",
                        "k8s.resourcequota.name"
                    ],
                    "timeAggregation": "DEFAULT"
                }
            ],
            "queriesSettings": {
                "resolution": ""
            },
            "tileFilter": {
                "timeframe": "-5m"
            },
            "tileType": "DATA_EXPLORER",
            "visualConfig": {
                "axes": {
                    "xAxis": {
                        "visible": true
                    },
                    "yAxes": []
                },
                "global": {
                    "hideLegend": false
                },
                "graphChartSettings": {
                    "connectNulls": false
                },
                "heatmapSettings": {
                    "yAxis": "VALUE"
                },
                "honeycombSettings": {
                    "showHive": true,
                    "showLabels": false,
                    "showLegend": true
                },
                "rules": [
                    {
                        "matcher": "B:",
                        "properties": {
                            "color": "DEFAULT"
                        },
                        "seriesOverrides": []
                    }
                ],
                "tableSettings": {
                    "isThresholdBackgroundAppliedToCell": false
                },
                "thresholds": [
                    {
                        "axisTarget": "LEFT",
                        "queryId": "",
                        "rules": [
                            {
                                "color": "#7dc540"
                            },
                            {
                                "color": "#f5d30f"
                            },
                            {
                                "color": "#dc172a"
                            }
                        ],
                        "visible": true
                    }
                ],
                "type": "TOP_LIST"
            }
        },
        {
            "assignedEntities": [],
            "bounds": {
                "height": 304,
                "left": 304,
                "top": 76,
                "width": 532
            },
            "chartVisible": false,
            "configured": true,
            "customName": "CPU requests quota used",
            "excludeMaintenanceWindows": false,
            "metricExpressions": [
                "resolution=null\u0026((builtin:kubernetes.resourcequota.requests_cpu_used/builtin:kubernetes.resourcequota.requests_cpu*100):splitBy(\"dt.entity.cloud_application_namespace\",\"k8s.resourcequota.name\"):sort(value(avg,descending)):setUnit(Percent):limit(10)):limit(100):names"
            ],
            "name": "Top 10 quota used in %",
            "queries": [
                {
                    "enabled": true,
                    "id": "C",
                    "metricSelector": "(builtin:kubernetes.resourcequota.requests_cpu_used / builtin:kubernetes.resourcequota.requests_cpu * 100):splitBy(\"dt.entity.cloud_application_namespace\",\"k8s.resourcequota.name\"):sort(value(avg,descending)):setUnit(Percent):limit(10)",
                    "splitBy": [
                        "dt.entity.cloud_application_namespace",
                        "k8s.resourcequota.name"
                    ],
                    "timeAggregation": "DEFAULT"
                }
            ],
            "queriesSettings": {
                "resolution": ""
            },
            "tileType": "DATA_EXPLORER",
            "visualConfig": {
                "axes": {
                    "xAxis": {
                        "displayName": "",
                        "visible": true
                    },
                    "yAxes": [
                        {
                            "defaultAxis": true,
                            "displayName": "",
                            "max": "100",
                            "min": "0",
                            "position": "LEFT",
                            "queryIds": [
                                "C"
                            ],
                            "visible": true
                        }
                    ]
                },
                "global": {
                    "hideLegend": false
                },
                "graphChartSettings": {
                    "connectNulls": false
                },
                "heatmapSettings": {
                    "yAxis": "VALUE"
                },
                "honeycombSettings": {
                    "showHive": true,
                    "showLabels": false,
                    "showLegend": true
                },
                "rules": [
                    {
                        "matcher": "C:",
                        "properties": {
                            "color": "DEFAULT",
                            "seriesType": "LINE"
                        },
                        "seriesOverrides": [],
                        "unitTransform": "Percent",
                        "valueFormat": "none"
                    }
                ],
                "tableSettings": {
                    "isThresholdBackgroundAppliedToCell": false
                },
                "thresholds": [
                    {
                        "axisTarget": "LEFT",
                        "queryId": "",
                        "rules": [
                            {
                                "color": "#7dc540"
                            },
                            {
                                "color": "#f5d30f"
                            },
                            {
                                "color": "#f5d30f",
                                "value": 90
                            }
                        ],
                        "visible": true
                    }
                ],
                "type": "GRAPH_CHART"
            }
        },
        {
            "assignedEntities": [],
            "bounds": {
                "height": 304,
                "left": 836,
                "top": 76,
                "width": 304
            },
            "chartVisible": false,
            "configured": true,
            "customName": "CPU limits quota used",
            "excludeMaintenanceWindows": false,
            "metricExpressions": [
                "resolution=null\u0026(builtin:kubernetes.resourcequota.limits_cpu_used:last:splitBy(\"dt.entity.cloud_application_namespace\",\"k8s.resourcequota.name\"):sum:sort(value(sum,descending))):limit(100):names:fold(auto)"
            ],
            "name": "Quota used",
            "queries": [
                {
                    "enabled": true,
                    "id": "B",
                    "metricSelector": "builtin:kubernetes.resourcequota.limits_cpu_used:last:splitBy(\"dt.entity.cloud_application_namespace\",\"k8s.resourcequota.name\"):sum:sort(value(sum,descending))",
                    "splitBy": [
                        "dt.entity.cloud_application_namespace",
                        "k8s.resourcequota.name"
                    ],
                    "timeAggregation": "DEFAULT"
                }
            ],
            "queriesSettings": {
                "resolution": ""
            },
            "tileFilter": {
                "timeframe": "-5m"
            },
            "tileType": "DATA_EXPLORER",
            "visualConfig": {
                "axes": {
                    "xAxis": {
                        "visible": true
                    },
                    "yAxes": []
                },
                "global": {
                    "hideLegend": false
                },
                "graphChartSettings": {
                    "connectNulls": false
                },
                "heatmapSettings": {
                    "yAxis": "VALUE"
                },
                "honeycombSettings": {
                    "showHive": true,
                    "showLabels": false,
                    "showLegend": true
                },
                "rules": [
                    {
                        "matcher": "B:",
                        "properties": {
                            "color": "DEFAULT"
                        },
                        "seriesOverrides": []
                    }
                ],
                "tableSettings": {
                    "isThresholdBackgroundAppliedToCell": false
                },
                "thresholds": [
                    {
                        "axisTarget": "LEFT",
                        "queryId": "",
                        "rules": [
                            {
                                "color": "#7dc540"
                            },
                            {
                                "color": "#f5d30f"
                            },
                            {
                                "color": "#dc172a"
                            }
                        ],
                        "visible": true
                    }
                ],
                "type": "TOP_LIST"
            }
        },
        {
            "assignedEntities": [],
            "bounds": {
                "height": 304,
                "left": 1140,
                "top": 76,
                "width": 532
            },
            "chartVisible": false,
            "configured": true,
            "customName": "CPU limits quota used",
            "excludeMaintenanceWindows": false,
            "metricExpressions": [
                "resolution=null\u0026((builtin:kubernetes.resourcequota.limits_cpu_used/builtin:kubernetes.resourcequota.limits_cpu*100):splitBy(\"dt.entity.cloud_application_namespace\",\"k8s.resourcequota.name\"):sort(value(avg,descending)):setUnit(Percent):limit(10)):limit(100):names"
            ],
            "name": "Top 10 quota used in %",
            "queries": [
                {
                    "enabled": true,
                    "id": "C",
                    "metricSelector": "(builtin:kubernetes.resourcequota.limits_cpu_used / builtin:kubernetes.resourcequota.limits_cpu * 100):splitBy(\"dt.entity.cloud_application_namespace\",\"k8s.resourcequota.name\"):sort(value(avg,descending)):setUnit(Percent):limit(10)",
                    "splitBy": [
                        "dt.entity.cloud_application_namespace",
                        "k8s.resourcequota.name"
                    ],
                    "timeAggregation": "DEFAULT"
                }
            ],
            "queriesSettings": {
                "resolution": ""
            },
            "tileType": "DATA_EXPLORER",
            "visualConfig": {
                "axes": {
                    "xAxis": {
                        "displayName": "",
                        "visible": true
                    },
                    "yAxes": [
                        {
                            "defaultAxis": true,
                            "displayName": "",
                            "max": "100",
                            "min": "0",
                            "position": "LEFT",
                            "queryIds": [
                                "C"
                            ],
                            "visible": true
                        }
                    ]
                },
                "global": {
                    "hideLegend": false
                },
                "graphChartSettings": {
                    "connectNulls": false
                },
                "heatmapSettings": {
                    "yAxis": "VALUE"
                },
                "honeycombSettings": {
                    "showHive": true,
                    "showLabels": false,
                    "showLegend": true
                },
                "rules": [
                    {
                        "matcher": "C:",
                        "properties": {
                            "color": "DEFAULT",
                            "seriesType": "LINE"
                        },
                        "seriesOverrides": [],
                        "unitTransform": "Percent",
                        "valueFormat": "auto"
                    }
                ],
                "tableSettings": {
                    "isThresholdBackgroundAppliedToCell": false
                },
                "thresholds": [
                    {
                        "axisTarget": "LEFT",
                        "queryId": "",
                        "rules": [
                            {
                                "color": "#7dc540"
                            },
                            {
                                "color": "#f5d30f"
                            },
                            {
                                "color": "#f5d30f",
                                "value": 90
                            }
                        ],
                        "visible": true
                    }
                ],
                "type": "GRAPH_CHART"
            }
        },
        {
            "assignedEntities": [],
            "bounds": {
                "height": 304,
                "left": 0,
                "top": 418,
                "width": 304
            },
            "chartVisible": false,
            "configured": true,
            "customName": "Memory requests quota used",
            "excludeMaintenanceWindows": false,
            "metricExpressions": [
                "resolution=null\u0026(builtin:kubernetes.resourcequota.requests_memory_used:last:splitBy(\"dt.entity.cloud_application_namespace\",\"k8s.resourcequota.name\"):sum:sort(value(sum,descending))):limit(100):names:fold(auto)"
            ],
            "name": "Quota used",
            "queries": [
                {
                    "enabled": true,
                    "id": "B",
                    "metricSelector": "builtin:kubernetes.resourcequota.requests_memory_used:last:splitBy(\"dt.entity.cloud_application_namespace\",\"k8s.resourcequota.name\"):sum:sort(value(sum,descending))",
                    "splitBy": [
                        "dt.entity.cloud_application_namespace",
                        "k8s.resourcequota.name"
                    ],
                    "timeAggregation": "DEFAULT"
                }
            ],
            "queriesSettings": {
                "resolution": ""
            },
            "tileFilter": {
                "timeframe": "-5m"
            },
            "tileType": "DATA_EXPLORER",
            "visualConfig": {
                "axes": {
                    "xAxis": {
                        "visible": true
                    },
                    "yAxes": []
                },
                "global": {
                    "hideLegend": false
                },
                "graphChartSettings": {
                    "connectNulls": false
                },
                "heatmapSettings": {
                    "yAxis": "VALUE"
                },
                "honeycombSettings": {
                    "showHive": true,
                    "showLabels": false,
                    "showLegend": true
                },
                "rules": [
                    {
                        "matcher": "B:",
                        "properties": {
                            "color": "DEFAULT"
                        },
                        "seriesOverrides": []
                    }
                ],
                "tableSettings": {
                    "isThresholdBackgroundAppliedToCell": false
                },
                "thresholds": [
                    {
                        "axisTarget": "LEFT",
                        "queryId": "",
                        "rules": [
                            {
                                "color": "#7dc540"
                            },
                            {
                                "color": "#f5d30f"
                            },
                            {
                                "color": "#dc172a"
                            }
                        ],
                        "visible": true
                    }
                ],
                "type": "TOP_LIST"
            }
        },
        {
            "assignedEntities": [],
            "bounds": {
                "height": 304,
                "left": 304,
                "top": 418,
                "width": 532
            },
            "chartVisible": false,
            "configured": true,
            "customName": "Memory requests quota used",
            "excludeMaintenanceWindows": false,
            "metricExpressions": [
                "resolution=null\u0026((builtin:kubernetes.resourcequota.requests_memory_used/builtin:kubernetes.resourcequota.requests_memory*100):splitBy(\"dt.entity.cloud_application_namespace\",\"k8s.resourcequota.name\"):sort(value(avg,descending)):setUnit(Percent):limit(10)):limit(100):names"
            ],
            "name": "Top 10 quota used in %",
            "queries": [
                {
                    "enabled": true,
                    "id": "B",
                    "metricSelector": "(builtin:kubernetes.resourcequota.requests_memory_used / builtin:kubernetes.resourcequota.requests_memory * 100):splitBy(\"dt.entity.cloud_application_namespace\",\"k8s.resourcequota.name\"):sort(value(avg,descending)):setUnit(Percent):limit(10)",
                    "splitBy": [
                        "dt.entity.cloud_application_namespace",
                        "k8s.resourcequota.name"
                    ],
                    "timeAggregation": "DEFAULT"
                }
            ],
            "queriesSettings": {
                "resolution": ""
            },
            "tileType": "DATA_EXPLORER",
            "visualConfig": {
                "axes": {
                    "xAxis": {
                        "displayName": "",
                        "visible": true
                    },
                    "yAxes": [
                        {
                            "defaultAxis": true,
                            "displayName": "",
                            "max": "100",
                            "min": "0",
                            "position": "LEFT",
                            "queryIds": [
                                "B"
                            ],
                            "visible": true
                        }
                    ]
                },
                "global": {
                    "hideLegend": false
                },
                "graphChartSettings": {
                    "connectNulls": false
                },
                "heatmapSettings": {
                    "yAxis": "VALUE"
                },
                "honeycombSettings": {
                    "showHive": true,
                    "showLabels": false,
                    "showLegend": true
                },
                "rules": [
                    {
                        "matcher": "B:",
                        "properties": {
                            "color": "DEFAULT",
                            "seriesType": "LINE"
                        },
                        "seriesOverrides": [],
                        "unitTransform": "Percent",
                        "valueFormat": "auto"
                    }
                ],
                "tableSettings": {
                    "isThresholdBackgroundAppliedToCell": false
                },
                "thresholds": [
                    {
                        "axisTarget": "LEFT",
                        "queryId": "",
                        "rules": [
                            {
                                "color": "#7dc540"
                            },
                            {
                                "color": "#f5d30f"
                            },
                            {
                                "color": "#f5d30f",
                                "value": 90
                            }
                        ],
                        "visible": true
                    }
                ],
                "type": "GRAPH_CHART"
            }
        },
        {
            "assignedEntities": [],
            "bounds": {
                "height": 304,
                "left": 836,
                "top": 418,
                "width": 304
            },
            "chartVisible": false,
            "configured": true,
            "customName": "Memory limits quota used",
            "excludeMaintenanceWindows": false,
            "metricExpressions": [
                "resolution=null\u0026(builtin:kubernetes.resourcequota.limits_memory_used:last:splitBy(\"dt.entity.cloud_application_namespace\",\"k8s.resourcequota.name\"):sum:sort(value(sum,descending))):limit(100):names:fold(auto)"
            ],
            "name": "Quota used",
            "queries": [
                {
                    "enabled": true,
                    "id": "B",
                    "metricSelector": "builtin:kubernetes.resourcequota.limits_memory_used:last:splitBy(\"dt.entity.cloud_application_namespace\",\"k8s.resourcequota.name\"):sum:sort(value(sum,descending))",
                    "splitBy": [
                        "dt.entity.cloud_application_namespace",
                        "k8s.resourcequota.name"
                    ],
                    "timeAggregation": "DEFAULT"
                }
            ],
            "queriesSettings": {
                "resolution": ""
            },
            "tileFilter": {
                "timeframe": "-5m"
            },
            "tileType": "DATA_EXPLORER",
            "visualConfig": {
                "axes": {
                    "xAxis": {
                        "visible": true
                    },
                    "yAxes": []
                },
                "global": {
                    "hideLegend": false
                },
                "graphChartSettings": {
                    "connectNulls": false
                },
                "heatmapSettings": {
                    "yAxis": "VALUE"
                },
                "honeycombSettings": {
                    "showHive": true,
                    "showLabels": false,
                    "showLegend": true
                },
                "rules": [
                    {
                        "matcher": "B:",
                        "properties": {
                            "color": "DEFAULT"
                        },
                        "seriesOverrides": []
                    }
                ],
                "tableSettings": {
                    "isThresholdBackgroundAppliedToCell": false
                },
                "thresholds": [
                    {
                        "axisTarget": "LEFT",
                        "queryId": "",
                        "rules": [
                            {
                                "color": "#7dc540"
                            },
                            {
                                "color": "#f5d30f"
                            },
                            {
                                "color": "#dc172a"
                            }
                        ],
                        "visible": true
                    }
                ],
                "type": "TOP_LIST"
            }
        },
        {
            "assignedEntities": [],
            "bounds": {
                "height": 304,
                "left": 1140,
                "top": 418,
                "width": 532
            },
            "chartVisible": false,
            "configured": true,
            "customName": "Memory limits quota used",
            "excludeMaintenanceWindows": false,
            "metricExpressions": [
                "resolution=null\u0026((builtin:kubernetes.resourcequota.limits_memory_used/builtin:kubernetes.resourcequota.limits_memory*100):splitBy(\"dt.entity.cloud_application_namespace\",\"k8s.resourcequota.name\"):sort(value(avg,descending)):setUnit(Percent):limit(10)):limit(100):names"
            ],
            "name": "Top 10 quota used in %",
            "queries": [
                {
                    "enabled": true,
                    "id": "B",
                    "metricSelector": "(builtin:kubernetes.resourcequota.limits_memory_used / builtin:kubernetes.resourcequota.limits_memory * 100):splitBy(\"dt.entity.cloud_application_namespace\",\"k8s.resourcequota.name\"):sort(value(avg,descending)):setUnit(Percent):limit(10)",
                    "splitBy": [
                        "dt.entity.cloud_application_namespace",
                        "k8s.resourcequota.name"
                    ],
                    "timeAggregation": "DEFAULT"
                }
            ],
            "queriesSettings": {
                "resolution": ""
            },
            "tileType": "DATA_EXPLORER",
            "visualConfig": {
                "axes": {
                    "xAxis": {
                        "displayName": "",
                        "visible": true
                    },
                    "yAxes": [
                        {
                            "defaultAxis": true,
                            "displayName": "",
                            "max": "100",
                            "min": "0",
                            "position": "LEFT",
                            "queryIds": [
                                "B"
                            ],
                            "visible": true
                        }
                    ]
                },
                "global": {
                    "hideLegend": false
                },
                "graphChartSettings": {
                    "connectNulls": false
                },
                "heatmapSettings": {
                    "yAxis": "VALUE"
                },
                "honeycombSettings": {
                    "showHive": true,
                    "showLabels": false,
                    "showLegend": true
                },
                "rules": [
                    {
                        "matcher": "B:",
                        "properties": {
                            "color": "DEFAULT",
                            "seriesType": "LINE"
                        },
                        "seriesOverrides": [],
                        "unitTransform": "Percent",
                        "valueFormat": "auto"
                    }
                ],
                "tableSettings": {
                    "isThresholdBackgroundAppliedToCell": false
                },
                "thresholds": [
                    {
                        "axisTarget": "LEFT",
                        "queryId": "",
                        "rules": [
                            {
                                "color": "#7dc540"
                            },
                            {
                                "color": "#f5d30f"
                            },
                            {
                                "color": "#f5d30f",
                                "value": 90
                            }
                        ],
                        "visible": true
                    }
                ],
                "type": "GRAPH_CHART"
            }
        },
        {
            "assignedEntities": [],
            "bounds": {
                "height": 304,
                "left": 0,
                "top": 760,
                "width": 304
            },
            "chartVisible": false,
            "configured": true,
            "customName": "CPU requests quota used",
            "excludeMaintenanceWindows": false,
            "metricExpressions": [
                "resolution=null\u0026(builtin:kubernetes.resourcequota.pods_used:last:splitBy(\"dt.entity.cloud_application_namespace\",\"k8s.resourcequota.name\"):sum:sort(value(sum,descending))):limit(100):names:fold(auto)"
            ],
            "name": "Quota used",
            "queries": [
                {
                    "enabled": true,
                    "id": "B",
                    "metricSelector": "builtin:kubernetes.resourcequota.pods_used:last:splitBy(\"dt.entity.cloud_application_namespace\",\"k8s.resourcequota.name\"):sum:sort(value(sum,descending))",
                    "splitBy": [
                        "dt.entity.cloud_application_namespace",
                        "k8s.resourcequota.name"
                    ],
                    "timeAggregation": "DEFAULT"
                }
            ],
            "queriesSettings": {
                "resolution": ""
            },
            "tileFilter": {
                "timeframe": "-5m"
            },
            "tileType": "DATA_EXPLORER",
            "visualConfig": {
                "axes": {
                    "xAxis": {
                        "visible": true
                    },
                    "yAxes": []
                },
                "global": {
                    "hideLegend": false
                },
                "graphChartSettings": {
                    "connectNulls": false
                },
                "heatmapSettings": {
                    "yAxis": "VALUE"
                },
                "honeycombSettings": {
                    "showHive": true,
                    "showLabels": false,
                    "showLegend": true
                },
                "rules": [
                    {
                        "matcher": "B:",
                        "properties": {
                            "color": "DEFAULT"
                        },
                        "seriesOverrides": []
                    }
                ],
                "tableSettings": {
                    "isThresholdBackgroundAppliedToCell": false
                },
                "thresholds": [
                    {
                        "axisTarget": "LEFT",
                        "queryId": "",
                        "rules": [
                            {
                                "color": "#7dc540"
                            },
                            {
                                "color": "#f5d30f"
                            },
                            {
                                "color": "#dc172a"
                            }
                        ],
                        "visible": true
                    }
                ],
                "type": "TOP_LIST"
            }
        },
        {
            "assignedEntities": [],
            "bounds": {
                "height": 304,
                "left": 304,
                "top": 760,
                "width": 532
            },
            "chartVisible": false,
            "configured": true,
            "customName": "CPU requests quota used",
            "excludeMaintenanceWindows": false,
            "metricExpressions": [
                "resolution=null\u0026((builtin:kubernetes.resourcequota.pods_used/builtin:kubernetes.resourcequota.pods*100):splitBy(\"dt.entity.cloud_application_namespace\",\"k8s.resourcequota.name\"):sort(value(avg,descending)):setUnit(Percent):limit(10)):limit(100):names"
            ],
            "name": "Top 10 quota used in %",
            "queries": [
                {
                    "enabled": true,
                    "id": "C",
                    "metricSelector": "(builtin:kubernetes.resourcequota.pods_used / builtin:kubernetes.resourcequota.pods * 100):splitBy(\"dt.entity.cloud_application_namespace\",\"k8s.resourcequota.name\"):sort(value(avg,descending)):setUnit(Percent):limit(10)",
                    "splitBy": [
                        "dt.entity.cloud_application_namespace",
                        "k8s.resourcequota.name"
                    ],
                    "timeAggregation": "DEFAULT"
                }
            ],
            "queriesSettings": {
                "resolution": ""
            },
            "tileType": "DATA_EXPLORER",
            "visualConfig": {
                "axes": {
                    "xAxis": {
                        "displayName": "",
                        "visible": true
                    },
                    "yAxes": [
                        {
                            "defaultAxis": true,
                            "displayName": "",
                            "max": "100",
                            "min": "0",
                            "position": "LEFT",
                            "queryIds": [
                                "C"
                            ],
                            "visible": true
                        }
                    ]
                },
                "global": {
                    "hideLegend": false
                },
                "graphChartSettings": {
                    "connectNulls": false
                },
                "heatmapSettings": {
                    "yAxis": "VALUE"
                },
                "honeycombSettings": {
                    "showHive": true,
                    "showLabels": false,
                    "showLegend": true
                },
                "rules": [
                    {
                        "matcher": "C:",
                        "properties": {
                            "color": "DEFAULT",
                            "seriesType": "LINE"
                        },
                        "seriesOverrides": [],
                        "unitTransform": "Percent",
                        "valueFormat": "none"
                    }
                ],
                "tableSettings": {
                    "isThresholdBackgroundAppliedToCell": false
                },
                "thresholds": [
                    {
                        "axisTarget": "LEFT",
                        "queryId": "",
                        "rules": [
                            {
                                "color": "#7dc540"
                            },
                            {
                                "color": "#f5d30f"
                            },
                            {
                                "color": "#f5d30f",
                                "value": 90
                            }
                        ],
                        "visible": true
                    }
                ],
                "type": "GRAPH_CHART"
            }
        }
    ]
}`
