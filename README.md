# Simple CICD with Jenkins to Docker service
Jenkins adalah tools untuk mengotomatisasi proses pembuatan, pengujian, dan penerapan software.
Saat melakukan setup ini author menggunakan server dengan OS type *Ubuntu 22.04 LTS*, dengan 2 core 8 Gb RAM, dan 50Gb disk.
## 1. Setup Jenkins server
```bash
sudo su

sudo wget -O /usr/share/keyrings/jenkins-keyring.asc \
  https://pkg.jenkins.io/debian-stable/jenkins.io-2023.key

echo "deb [signed-by=/usr/share/keyrings/jenkins-keyring.asc]" \
  https://pkg.jenkins.io/debian-stable binary/ | sudo tee \
  /etc/apt/sources.list.d/jenkins.list > /dev/null

sudo apt-get update
sudo apt-get install jenkins
```
By default jenkins berjalan di port 8080, dan diakses melalui http pada :
> http://localhost:8080/

Admin credential, dapat ditemukan dengan 
```bash
cat /var/lib/jenkins/secrets/initialAdminPassword
```

Karena pada projek ini output akan berupa docker container, maka server perlu diinstalkan docker juga.
```bash
curl -fsSL https://get.docker.com -o get-docker.sh
sh get-docker.sh

# Tambahkan user jenkins ke dalam Docker grup
usermod -aG docker jenkins
```
Jenkins membutuhkan authorization agar dapat mendeploy aplikasi di server, maka dibuthkan authorize key
```bash
ssh-keygen
cat /root/.ssh/id_rsa.pub
```
Copy output command di atas, key tersebut dibutuhkan oleh jenkins agar dapat mengakses ke server.
(Optional jika ditambahkan SSH Authorize pada Github/Gitlab, tapi tidak direkomendasikan)
## 2. Setup Git repository
Buat repository baru ke git, dan push 3 file (app.js, Dockerfile, dan docker-compose.yaml) atau file projek apapun itu.

## 3. Build Job
Akses Jenkins dashboard di
> http://localhost:8080

Pada konfigurasi awal, pilih install recomended plugin, kemudian :
> 1. Create a job
> 2. Masukkan nama job dan pilih *Freestyle*
> 3. Optional untuk memasukkan deskripsi
> 4. Pada Source Code Management, pilih Git
>> 4.1 Klik Add, Jenkins

>> 4.2 Domain (Global Credential), Kind (SSH Username with private key), Scope (Global), Username (root), Enter directly, Add (paste id_rsa.pub yang telah di copy sebelumnya), Add

> 5. Pada Build Steps, pilih  *Execute sheel*, masukkan :
```bash
docker build -t myapp .
docker compose down
docker compose up -d
```
Command tersebut bertujuan agar setiap kali jenkins melakukan build, jenkins akan membuat image terbaru dari aplikasi yang di deploy. Kemudian akan stop dan remove container yang eksis dan membuat container baru dengan nama yang sama dan image terbaru. Hal ini dilakukan untuk menghindari konflik pada docker dengan nama container yang sama.