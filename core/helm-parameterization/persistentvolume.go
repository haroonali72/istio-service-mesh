package helm_parameterization

import (
	"errors"
	"istio-service-mesh/core/helm-parameterization/types"
	core "k8s.io/api/core/v1"
	"sigs.k8s.io/yaml"
	"strconv"
	"strings"
)

func PersistentVolumeParameters(persistentVolume *core.PersistentVolume) (persistentVolumeYaml []byte, persistentVolumeParams []byte, functionsData []byte, err error) {
	result, err := yaml.Marshal(persistentVolume)
	if err != nil {
		return nil, nil, nil, err
	}
	persistentVolumeRaw := new(types.PersistentVolumeTemplate)
	err = yaml.Unmarshal(result, persistentVolumeRaw)
	if err != nil {
		return nil, nil, nil, err
	}

	persistentVolumeParams = []byte("\n" + persistentVolume.Name + "PV:")
	if len(persistentVolume.Spec.MountOptions) > 0 {
		persistentVolumeParams = append(persistentVolumeParams, []byte("\n  mountOptions:")...)
		persistentVolumeRaw.Spec.MountOptions = "{{ index .Values \"" + persistentVolume.Name + "PV\" \"mountOptions\" | toYaml | trim | nindent 2 }}"

	}
	for _, each := range persistentVolume.Spec.MountOptions {
		persistentVolumeParams = append(persistentVolumeParams, []byte("\n    - "+each)...)
	}
	//	persistentVolumeClaimRaw.Spec.AccessModes="{{ index .Values \""+ persistentVolumeClaim.Name+"PVC\" \"accessModes\" | toYaml | trim | nindent 2 }}"
	qunatity, ok := persistentVolume.Spec.Capacity["storage"]
	if ok {
		persistentVolumeParams = append(persistentVolumeParams, []byte("\n  capacity: "+qunatity.String())...)
		persistentVolumeRaw.Spec.Capacity["storage"] = "{{ index .Values \"" + persistentVolume.Name + "PV\" \"capacity\" }}"
		//	persistentVolumeRaw.Spec.Capacity["storage"]="{{ .Values."+ persistentVolume.Name+"PV.capacity.storage }}"
	}
	if persistentVolume.Spec.VolumeMode != nil {
		persistentVolumeParams = append(persistentVolumeParams, []byte("\n  volumeMode: "+*persistentVolume.Spec.VolumeMode)...)
		//		persistentVolumeRaw.Spec.VolumeMode="{{ .Values."+ persistentVolume.Name+"PV.volumeMode }}"
		persistentVolumeRaw.Spec.VolumeMode = "{{ index .Values \"" + persistentVolume.Name + "PV\" \"volumeMode\" }}"

	}
	if len(persistentVolume.Spec.AccessModes) > 0 {
		persistentVolumeRaw.Spec.AccessModes = "{{ index .Values \"" + persistentVolume.Name + "PV\" \"accessModes\" | toYaml | trim | nindent 2 }}"
		persistentVolumeParams = append(persistentVolumeParams, []byte("\n  accessModes:")...)
	}

	for _, each := range persistentVolume.Spec.AccessModes {
		persistentVolumeParams = append(persistentVolumeParams, []byte("\n    - "+each)...)
	}
	//persistentVolumeRaw.Spec.MountOptions="{{ .Values."+ persistentVolume.Name+"PV.accessModes }}"

	if persistentVolume.Spec.StorageClassName != "" {
		persistentVolumeParams = append(persistentVolumeParams, []byte("\n  storageClassName: "+persistentVolume.Spec.StorageClassName)...)
		//persistentVolumeRaw.Spec.StorageClassName="{{ .Values."+ persistentVolume.Name+"PV.storageClassName }}"
		persistentVolumeRaw.Spec.StorageClassName = "{{ index .Values \"" + persistentVolume.Name + "PV\" \"storageClassName\" }}"

	}
	if persistentVolume.Spec.PersistentVolumeReclaimPolicy != "" {
		persistentVolumeParams = append(persistentVolumeParams, []byte("\n  persistentVolumeReclaimPolicy: "+persistentVolume.Spec.StorageClassName)...)
		//	persistentVolumeRaw.Spec.PersistentVolumeReclaimPolicy="{{ .Values."+ persistentVolume.Name+"PV.persistentVolumeReclaimPolicy }}"
		persistentVolumeRaw.Spec.PersistentVolumeReclaimPolicy = "{{ index .Values \"" + persistentVolume.Name + "PV\" \"persistentVolumeReclaimPolicy\" }}"
	}
	if persistentVolume.Spec.GCEPersistentDisk != nil {
		persistentVolumeParams = append(persistentVolumeParams, []byte("\n  gcePersistentDisk:")...)
		if persistentVolume.Spec.GCEPersistentDisk.PDName == "" {
			return nil, nil, nil, errors.New("can not find pdName in gcePersistentDisk")
		}

		persistentVolumeParams = append(persistentVolumeParams, []byte("\n    pdName: "+persistentVolume.Spec.GCEPersistentDisk.PDName)...)
		//		persistentVolumeRaw.Spec.GCEPersistentDisk.PdName="{{ .Values."+ persistentVolume.Name+"PV.gcePersistentDisk.pdName }}"
		persistentVolumeRaw.Spec.GCEPersistentDisk.PdName = "{{ index .Values \"" + persistentVolume.Name + "PV\" \"gcePersistentDisk\" \"pdName\" }}"
		if persistentVolume.Spec.GCEPersistentDisk.FSType != "" {
			persistentVolumeParams = append(persistentVolumeParams, []byte("\n    fsType: "+persistentVolume.Spec.GCEPersistentDisk.FSType)...)
			//persistentVolumeRaw.Spec.GCEPersistentDisk.FSType="{{ .Values."+ persistentVolume.Name+"PV.gcePersistentDisk.fsType }}"
			persistentVolumeRaw.Spec.GCEPersistentDisk.FSType = "{{ index .Values \"" + persistentVolume.Name + "PV\" \"gcePersistentDisk\" \"fsType\" }}"
		}
		persistentVolumeParams = append(persistentVolumeParams, []byte("\n    readOnly: "+strconv.FormatBool(persistentVolume.Spec.GCEPersistentDisk.ReadOnly))...)
		//			persistentVolumeRaw.Spec.GCEPersistentDisk.ReadOnly="{{ .Values."+ persistentVolume.Name+"PV.gcePersistentDisk.readOnly }}"
		persistentVolumeRaw.Spec.GCEPersistentDisk.ReadOnly = "{{ index .Values \"" + persistentVolume.Name + "PV\" \"gcePersistentDisk\" \"readOnly\" }}"
		persistentVolumeParams = append(persistentVolumeParams, []byte("\n    partition: "+strconv.FormatInt(int64(persistentVolume.Spec.GCEPersistentDisk.Partition), 10))...)
		//		persistentVolumeRaw.Spec.GCEPersistentDisk.Partition="{{ .Values."+ persistentVolume.Name+"PV.gcePersistentDisk.partition }}"
		persistentVolumeRaw.Spec.GCEPersistentDisk.Partition = "{{ index .Values \"" + persistentVolume.Name + "PV\" \"gcePersistentDisk\" \"partition\" }}"
	} else if persistentVolume.Spec.AWSElasticBlockStore != nil {
		persistentVolumeParams = append(persistentVolumeParams, []byte("\n  awsElasticBlockStore:")...)
		if persistentVolume.Spec.AWSElasticBlockStore.VolumeID == "" {
			return nil, nil, nil, errors.New("can not find volumeID in awsElasticBlockStore")
		}

		persistentVolumeParams = append(persistentVolumeParams, []byte("\n    volumeID: "+persistentVolume.Spec.AWSElasticBlockStore.VolumeID)...)
		//		persistentVolumeRaw.Spec.AWSElasticBlockStore.VolumeId="{{ .Values."+ persistentVolume.Name+"PV.awsElasticBlockStore.volumeID }}"
		persistentVolumeRaw.Spec.AWSElasticBlockStore.VolumeId = "{{ index .Values \"" + persistentVolume.Name + "PV\" \"awsElasticBlockStore\" \"volumeID\" }}"
		if persistentVolume.Spec.AWSElasticBlockStore.FSType != "" {
			persistentVolumeParams = append(persistentVolumeParams, []byte("\n    fsType: "+persistentVolume.Spec.AWSElasticBlockStore.FSType)...)
			//	persistentVolumeRaw.Spec.AWSElasticBlockStore.FSType="{{ .Values."+ persistentVolume.Name+"PV.awsElasticBlockStore.fsType }}"
			persistentVolumeRaw.Spec.AWSElasticBlockStore.FSType = "{{ index .Values \"" + persistentVolume.Name + "PV\" \"awsElasticBlockStore\" \"fsType\" }}"
		}
		persistentVolumeParams = append(persistentVolumeParams, []byte("\n    readOnly: "+strconv.FormatBool(persistentVolume.Spec.AWSElasticBlockStore.ReadOnly))...)
		//		persistentVolumeRaw.Spec.AWSElasticBlockStore.ReadOnly="{{ .Values."+ persistentVolume.Name+"PV.awsElasticBlockStore.readOnly }}"
		persistentVolumeRaw.Spec.AWSElasticBlockStore.ReadOnly = "{{ index .Values \"" + persistentVolume.Name + "PV\" \"awsElasticBlockStore\" \"readOnly\" }}"

		persistentVolumeParams = append(persistentVolumeParams, []byte("\n    partition: "+strconv.FormatInt(int64(persistentVolume.Spec.AWSElasticBlockStore.Partition), 10))...)
		//		persistentVolumeRaw.Spec.AWSElasticBlockStore.Partition="{{ .Values."+ persistentVolume.Name+"PV.awsElasticBlockStore.partition }}"
		persistentVolumeRaw.Spec.AWSElasticBlockStore.Partition = "{{ index .Values \"" + persistentVolume.Name + "PV\" \"awsElasticBlockStore\" \"partition\" }}"

	} else if persistentVolume.Spec.AzureDisk != nil {

		persistentVolumeParams = append(persistentVolumeParams, []byte("\n  azureDisk:")...)
		if persistentVolume.Spec.AzureDisk.DiskName == "" {
			return nil, nil, nil, errors.New("can not find diskName in azureDisk")
		}

		persistentVolumeParams = append(persistentVolumeParams, []byte("\n    diskName: "+persistentVolume.Spec.AzureDisk.DiskName)...)
		//persistentVolumeRaw.Spec.AzureDisk.DiskName="{{ .Values."+ persistentVolume.Name+"PV.azureDisk.diskName }}"
		persistentVolumeRaw.Spec.AzureDisk.DiskName = "{{ index .Values \"" + persistentVolume.Name + "PV\" \"azureDisk\" \"diskName\" }}"

		if persistentVolume.Spec.AzureDisk.DataDiskURI == "" {
			return nil, nil, nil, errors.New("can not find diskURI in azureDisk")
		}

		persistentVolumeParams = append(persistentVolumeParams, []byte("\n    diskURI: "+persistentVolume.Spec.AzureDisk.DataDiskURI)...)
		//		persistentVolumeRaw.Spec.AzureDisk.DiskURI="{{ .Values."+ persistentVolume.Name+"PV.azureDisk.diskURI }}"
		persistentVolumeRaw.Spec.AzureDisk.DiskURI = "{{ index .Values \"" + persistentVolume.Name + "PV\" \"azureDisk\" \"diskURI\" }}"
		if persistentVolume.Spec.AzureDisk.Kind == nil {
			return nil, nil, nil, errors.New("can not find kind in azureDisk")
		}

		persistentVolumeParams = append(persistentVolumeParams, []byte("\n    kind: "+persistentVolume.Spec.AzureDisk.DataDiskURI)...)
		//		persistentVolumeRaw.Spec.AzureDisk.Kind="{{ .Values."+ persistentVolume.Name+"PV.azureDisk.kind }}"
		persistentVolumeRaw.Spec.AzureDisk.Kind = "{{ index .Values \"" + persistentVolume.Name + "PV\" \"azureDisk\" \"kind\" }}"

		if persistentVolume.Spec.AzureDisk.FSType != nil {
			persistentVolumeParams = append(persistentVolumeParams, []byte("\n    fsType: "+*persistentVolume.Spec.AzureDisk.FSType)...)
			//		persistentVolumeRaw.Spec.AzureDisk.FSType="{{ .Values."+ persistentVolume.Name+"PV.azureDisk.fsType }}"
			persistentVolumeRaw.Spec.AzureDisk.FSType = "{{ index .Values \"" + persistentVolume.Name + "PV\" \"azureDisk\" \"fsType\" }}"
		}
		if persistentVolume.Spec.AzureDisk.ReadOnly != nil {
			persistentVolumeParams = append(persistentVolumeParams, []byte("\n    readOnly: "+strconv.FormatBool(*persistentVolume.Spec.AzureDisk.ReadOnly))...)
			//	persistentVolumeRaw.Spec.AzureDisk.ReadOnly="{{ .Values."+ persistentVolume.Name+"PV.azureDisk.readOnly }}"
			persistentVolumeRaw.Spec.AzureDisk.ReadOnly = "{{ index .Values \"" + persistentVolume.Name + "PV\" \"azureDisk\" \"readOnly\" }}"
		}
		if persistentVolume.Spec.AzureDisk.CachingMode != nil {
			persistentVolumeParams = append(persistentVolumeParams, []byte("\n    cachingMode: "+*persistentVolume.Spec.AzureDisk.CachingMode)...)
			//			persistentVolumeRaw.Spec.AzureDisk.CachingMode="{{ .Values."+ persistentVolume.Name+"PV.azureDisk.cachingMode }}"
			persistentVolumeRaw.Spec.AzureDisk.CachingMode = "{{ index .Values \"" + persistentVolume.Name + "PV\" \"azureDisk\" \"cachingMode\" }}"

		}

	} else if persistentVolume.Spec.AzureFile != nil {
		persistentVolumeParams = append(persistentVolumeParams, []byte("\n  azureFile:")...)
		if persistentVolume.Spec.AzureFile.SecretName == "" {
			return nil, nil, nil, errors.New("can not find secretName in azureFile")
		}

		persistentVolumeParams = append(persistentVolumeParams, []byte("\n    secretName: "+persistentVolume.Spec.AzureFile.SecretName)...)
		//	persistentVolumeRaw.Spec.AzureFile.SecretName="{{ .Values."+ persistentVolume.Name+"PV.azureFile.secretName }}"
		persistentVolumeRaw.Spec.AzureFile.SecretName = "{{ index .Values \"" + persistentVolume.Name + "PV\" \"azureFile\" \"secretName\" }}"

		if persistentVolume.Spec.AzureFile.ShareName == "" {
			return nil, nil, nil, errors.New("can not find shareName in azureFile")
		}

		persistentVolumeParams = append(persistentVolumeParams, []byte("\n    shareName: "+persistentVolume.Spec.AzureFile.ShareName)...)
		//		persistentVolumeRaw.Spec.AzureFile.ShareName="{{ .Values."+ persistentVolume.Name+"PV.azureFile.shareName }}"
		persistentVolumeRaw.Spec.AzureFile.ShareName = "{{ index .Values \"" + persistentVolume.Name + "PV\" \"azureFile\" \"shareName\" }}"

		if persistentVolume.Spec.AzureFile.SecretNamespace != nil {
			persistentVolumeParams = append(persistentVolumeParams, []byte("\n    secretNamespace: "+*persistentVolume.Spec.AzureFile.SecretNamespace)...)
			//			persistentVolumeRaw.Spec.AzureFile.SecretNamespace="{{ .Values."+ persistentVolume.Name+"PV.azureFile.secretNamespace }}"
			persistentVolumeRaw.Spec.AzureFile.SecretNamespace = "{{ index .Values \"" + persistentVolume.Name + "PV\" \"azureFile\" \"secretNamespace\" }}"

		}
		persistentVolumeParams = append(persistentVolumeParams, []byte("\n    readOnly: "+strconv.FormatBool(persistentVolume.Spec.AzureFile.ReadOnly))...)
		//persistentVolumeRaw.Spec.AzureFile.ReadOnly="{{ .Values."+ persistentVolume.Name+"PV.azureFile.readOnly }}"
		persistentVolumeRaw.Spec.AzureFile.ReadOnly = "{{ index .Values \"" + persistentVolume.Name + "PV\" \"azureFile\" \"readOnly\" }}"

	} else {
		return nil, nil, nil, errors.New("unsupported persistentVolume source")
	}

	persistentVolumeYaml, err = yaml.Marshal(persistentVolumeRaw)
	temp := strings.ReplaceAll(string(persistentVolumeYaml), "'{{", "{{")
	temp = strings.ReplaceAll(temp, "}}'", "}}")
	persistentVolumeYaml = []byte(temp)
	return
}
