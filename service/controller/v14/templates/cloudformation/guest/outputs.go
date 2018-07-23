package guest

const Outputs = `{{define "outputs"}}
Outputs:
  DockerVolumeResourceName:
    Value: {{ .Outputs.Master.DockerVolume.ResourceName }}
  {{ if .Route53Enabled }}
  HostedZoneNameServers:
    Value: !Join [ ',', !GetAtt 'HostedZone.NameServers' ]
  {{ end }}
  MasterImageID:
    Value: {{ .Outputs.Master.ImageID }}
  MasterInstanceResourceName:
    Value: {{ .Outputs.Master.Instance.ResourceName }}
  MasterInstanceType:
    Value: {{ .Outputs.Master.Instance.Type }}
  MasterCloudConfigVersion:
    Value: {{ .Outputs.Master.CloudConfig.Version }}
  {{ .Outputs.Worker.ASG.Key }}:
    Value: !Ref {{ .Outputs.Worker.ASG.Ref }}
  WorkerCount:
    Value: {{ .Outputs.Worker.Count }}
  WorkerImageID:
    Value: {{ .Outputs.Worker.ImageID }}
  WorkerInstanceType:
    Value: {{ .Outputs.Worker.InstanceType }}
  WorkerCloudConfigVersion:
    Value: {{ .Outputs.Worker.CloudConfig.Version }}
  VersionBundleVersion:
    Value:
      Ref: VersionBundleVersionParameter
{{end}}`