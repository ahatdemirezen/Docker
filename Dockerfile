FROM golang:1.22.2-alpine
RUN mkdir app
# run mkdir app = çalışılacak dizede  build adında klasör oluşturur veya  çalışma dizini oluşturur
WORKDIR /app  
# workdır /app = dockerfilede belirtilen containerin çalışma dizinini belirler çalışma dizinini app olarak belirler

COPY . .
# ilk nokta dockerfile içinde bulunduğu tüm dosyaları ifade eder onları ele alır 
# ikinci nokta gönderileceği yeri temsil eder yani app 'i '
RUN go build -o ascii-art-dockerize . 
# bu kısımda  derleme işlemi yapılır -o isminin verildiği kısım ascii.. adında go build ise = derleme komutu
ENTRYPOINT [ "./ascii-art-dockerize" ]
# entrypoınt çalıştıralacağını belirtir . = çalışma dizinini belirtir. asciiart.. = dosyasının yürütülmesini sağlar.
EXPOSE 8080
# çalışacapı portal        