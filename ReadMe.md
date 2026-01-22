## Task 1 :
ArticatHub üzerinden resmi nginx char'ının K3S üzerine kurulumu
<hr>

### Görev Adımları
1) Bitnami repo'sunun helm'e eklenmesi
```bash
helm repo add bitnami https://charts.bitnami.com/bitnami
helm repo update
```

2) bitnami/nginx chart'ının kurulumu
```bash
helm install my-nginx bitnami/nginx --version 22.4.3
```

3) Kurulumların doğrulanması
```bash
helm ls #release listesi
kubectl get pods 
```

4) Helm Status çıktısının incelenmesi
```bash
helm status my-nginx
```
```bash
NAME: my-nginx
LAST DEPLOYED: Wed Jan 21 13:58:21 2026
NAMESPACE: default
STATUS: deployed
REVISION: 1
TEST SUITE: None
NOTES:
CHART NAME: nginx
CHART VERSION: 22.4.3
APP VERSION: 1.29.4

⚠ WARNING: Since August 28th, 2025, only a limited subset of images/charts are available for free.
    Subscribe to Bitnami Secure Images to receive continued support and security updates.
    More info at https://bitnami.com and https://github.com/bitnami/containers/issues/83267

** Please be patient while the chart is being deployed **
NGINX can be accessed through the following DNS name from within your cluster:

    my-nginx.default.svc.cluster.local (port 80)
```

#### Soru : Buradaki NOTES kısmı ne işe yarar?
- uygulamaya nasıl erişim sağlanacağı hakkında bilgi
- deployment süreci hakkında bilgi
- önemli uyarılar
- chart’a özel best practice önerileri, bulundurur.

5) Port Forward ile localden erişim
```bash
kubectl edit svc my-nginx
```
```bash
 ports:
  - name: http
    nodePort: 30080
    port: 80
    protocol: TCP
    targetPort: http
```
> Vagrantfile'ın düzenlenmesi
```bash
config.vm.network "forwarded_port", guest: 30080, host: 8080, host_ip: "127.0.0.1"
```
> vagrant reload işlemi
```bash
vagrant reload
```

![port-forward-check](ss/port-forward.png)

## Task 2 :
Kurulu bir uygulamanın özelliklerini dışarıdan müdahale ederek değiştirmek.
<hr>

### Görev Adımları
1) mevcut value'ları inceleme
```bash
helm show values bitnami/nginx > default-values.yaml
```
2) Yaml dosyasını düzenleme
> görev kapsamında bizde istenen replica sayısını 1'den 3'e çıkarmak, uygulamanın 1 CPU 2 RAM ile çalışmasını sağlamak ancak vm'in sahip olduğu donanımsal kaynaklardan dolayı cpu ve ram miktarı şu şekilde düzenlenmiştir : 0.2 CPU - 128Mi RAM

![replica-count](ss/replica-count.png)
![resources](ss/resources.png)

> Resorces alanında yer alan *resourcesPreset : nano*, bitnami chart'larında bu ayar, kaynakları otomatik olarak çok düşük (nano) seviyeye çeker. Yaptığımız değişiklikleri ezmemesi adına bu "none" olarak güncellememiz gerekir.

3) Değişikliklerin uygulanması
```bash
helm upgrade my-nginx bitnami/nginx -f values.yaml 
```
> upgrade çıktısı
```bash
vagrant@vagrant-ubuntu-jammy:~$ helm upgrade my-nginx bitnami/nginx -f default-values.yaml
Release "my-nginx" has been upgraded. Happy Helming!
NAME: my-nginx
LAST DEPLOYED: Thu Jan 22 14:05:07 2026
NAMESPACE: default
STATUS: deployed
REVISION: 2
TEST SUITE: None
NOTES:
CHART NAME: nginx
CHART VERSION: 22.4.3
APP VERSION: 1.29.4

⚠ WARNING: Since August 28th, 2025, only a limited subset of images/charts are available for free.
    Subscribe to Bitnami Secure Images to receive continued support and security updates.
    More info at https://bitnami.com and https://github.com/bitnami/containers/issues/83267

** Please be patient while the chart is being deployed **
NGINX can be accessed through the following DNS name from within your cluster:

    my-nginx.default.svc.cluster.local (port 80)
```
4) Güncel pod'ların incelenmesi
```bash
kubectl get pods --watch
```
![current pods](ss/current-pods.png)
