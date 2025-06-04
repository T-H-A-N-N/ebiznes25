# ebiznes25

Zadanie 1 Docker

✅ 3.0 obraz ubuntu z Pythonem w wersji 3.10

✅ 3.5 obraz ubuntu:24.02 z Javą w wersji 8 oraz Kotlinem

✅ 4.0 do powyższego należy dodać najnowszego Gradle’a oraz paczkę JDBC
SQLite w ramach projektu na Gradle (build.gradle) [commit](https://github.com/T-H-A-N-N/ebiznes25/commit/e5bb0bb19ae7c395f4dfef52fa064cae7a62e933)

4.5 stworzyć przykład typu HelloWorld oraz uruchomienie aplikacji
przez CMD oraz gradle

5.0 dodać konfigurację docker-compose

Obraz należy wysłać na hub.docker.com, a link do obrazu należy dodać w
README na githubie.

[docker hub](https://hub.docker.com/r/th4nn/ebiz1)

Zadanie 2 Scala

✅ 3.0 Należy stworzyć kontroler do Produktów

✅ 3.5 Do kontrolera należy stworzyć endpointy zgodnie z CRUD - dane
pobierane z listy

✅ 4.0 Należy stworzyć kontrolery do Kategorii oraz Koszyka + endpointy
zgodnie z CRUD [commit](https://github.com/T-H-A-N-N/ebiznes25/commit/19d54ffa5029ed7ae182f1871ff669bc37ae3b6d)

4.5 Należy aplikację uruchomić na dockerze (stworzyć obraz) oraz dodać
skrypt uruchamiający aplikację via ngrok

5.0 Należy dodać konfigurację CORS dla dwóch hostów dla metod CRUD

Kontrolery mogą bazować na listach zamiast baz danych. CRUD: show all,
show by id (get), update (put), delete (delete), add (post).

Zadanie 3 Kotlin

✅ 3.0 Należy stworzyć aplikację kliencką w Kotlinie we frameworku Ktor,
która pozwala na przesyłanie wiadomości na platformę Discord

✅ 3.5 Aplikacja jest w stanie odbierać wiadomości użytkowników z
platformy Discord skierowane do aplikacji (bota) [commit](https://github.com/T-H-A-N-N/ebiznes25/commit/d72d5adcb77048059d93796947ddca36f4e66fb8)

✅ 4.0 Zwróci listę kategorii na określone żądanie użytkownika

✅ 4.5 Zwróci listę produktów wg żądanej kategorii [commit](https://github.com/T-H-A-N-N/ebiznes25/commit/57a40eb2855848bc4ec6a59fb55a13cdbc1f1c1a)

✅ 5.0 Aplikacja obsłuży dodatkowo jedną z platform: Slack, Messenger,
Webex [commit](https://github.com/T-H-A-N-N/ebiznes25/commit/35d9713caaa60388f354d54dcf6e09e1a9ab6f99)

Zadanie 4 Go

Należy stworzyć projekt w echo w Go. Należy wykorzystać gorm do stworzenia 5 modeli, gdzie pomiędzy dwoma musi być relacja. Należy zaimplementować proste endpointy do dodawania oraz wyświetlania danych za pomocą modeli. Jako bazę danych można wybrać dowolną, sugerowałbym jednak pozostać przy sqlite.

✅ 3.0 Należy stworzyć aplikację we frameworki echo w j. Go, która będzie miała kontroler Produktów zgodny z CRUD [commit](https://github.com/T-H-A-N-N/ebiznes25/commit/5f532c037f31dc56e54b0cb9aa2ad97d3484aebd)

3.5 Należy stworzyć model Produktów wykorzystując gorm oraz wykorzystać model do obsługi produktów (CRUD) w kontrolerze (zamiast listy)

4.0 Należy dodać model Koszyka oraz dodać odpowiedni endpoint

4.5 Należy stworzyć model kategorii i dodać relację między kategorią, a produktem

5.0 pogrupować zapytania w gorm’owe scope'y

Zadanie 5 Frontend

Należy stworzyć aplikację kliencką wykorzystując bibliotekę React.js.
W ramach projektu należy stworzyć trzy komponenty: Produkty, Koszyk
oraz Płatności. Koszyk oraz Płatności powinny wysyłać do aplikacji
serwerowej dane, a w Produktach powinniśmy pobierać dane o produktach
z aplikacji serwerowej. Aplikacja serwera w jednym z trzech języków:
Kotlin, Scala, Go. Dane pomiędzy wszystkimi komponentami powinny być
przesyłane za pomocą React hooks.

✅ 3.0 W ramach projektu należy stworzyć dwa komponenty: Produkty oraz
Płatności; Płatności powinny wysyłać do aplikacji serwerowej dane, a w
Produktach powinniśmy pobierać dane o produktach z aplikacji
serwerowej; (nie chciało mi całej aplikacji na raz dodać więc musiałem kilka commitów) [commit1](https://github.com/T-H-A-N-N/ebiznes25/commit/96e67e0b1a42943f9df8add168a79dac2cec0fa1) 

3.5 Należy dodać Koszyk wraz z widokiem; należy wykorzystać routing

4.0 Dane pomiędzy wszystkimi komponentami powinny być przesyłane za
pomocą React hooks

4.5 Należy dodać skrypt uruchamiający aplikację serwerową oraz
kliencką na dockerze via docker-compose

5.0 Należy wykorzystać axios’a oraz dodać nagłówki pod CORS

Zadanie 6 testy

✅ 3.0 Należy stworzyć 20 przypadków testowych w CypressJS lub Selenium
(Kotlin, Python, Java, JS, Go, Scala)

✅ 3.5 Należy rozszerzyć testy funkcjonalne, aby zawierały minimum 50
asercji

✅ 4.0 Należy stworzyć testy jednostkowe do wybranego wcześniejszego
projektu z minimum 50 asercjami

✅ 4.5 Należy dodać testy API, należy pokryć wszystkie endpointy z
minimum jednym scenariuszem negatywnym per endpoint [commit](https://github.com/T-H-A-N-N/ebiznes25/commit/be89792763ea5fb43c4193baad1325f2857d9ce3)

✅ 5.0 Należy uruchomić testy funkcjonalne na Browserstacku [commit](https://github.com/T-H-A-N-N/ebiznes25/commit/f2f424ba0a01484aa67fb46e321b4205d4a038f5)

Zadanie 7 Sonar

✅ 3.0 Należy dodać litera do odpowiedniego kodu aplikacji serwerowej w
hookach gita [commit](https://github.com/T-H-A-N-N/ebiznes25/commit/fb408daa7709a37d22fe157351c0ab24f658ee01)

✅ 3.5 Należy wyeliminować wszystkie bugi w kodzie w Sonarze (kod
aplikacji serwerowej) [commit](https://github.com/T-H-A-N-N/ebiznes25/commit/6edaabc8eacb679db9e39bc7e4058de9fe324f4b)

4.0 Należy wyeliminować wszystkie zapaszki w kodzie w Sonarze (kod
aplikacji serwerowej)

4.5 Należy wyeliminować wszystkie podatności oraz błędy bezpieczeństwa
w kodzie w Sonarze (kod aplikacji serwerowej)

5.0 Należy wyeliminować wszystkie błędy oraz zapaszki w kodzie
aplikacji klienckiej

Zadanie 8 Oauth2

Należy skonfigurować klienta Oauth2 (4.0). Dane o użytkowniku wraz z
tokenem powinny być przechowywane po stronie bazy serwera, a nowy
token (inny niż ten od dostawcy) powinien zostać wysłany do klienta
(React). Można zastosować mechanizm sesji lub inny dowolny (5.0).
Zabronione jest tworzenie klientów bezpośrednio po stronie React'a
wyłączając z komunikacji aplikację serwerową, np. wykorzystując auth0.

Prawidłowa komunikacja: react-sewer-dostawca-serwer(via return
uri)-react.

✅ 3.0 logowanie przez aplikację serwerową (bez Oauth2) [commit1](https://github.com/T-H-A-N-N/ebiznes25/commit/5c88d56938869fb03688f01d1a6a74277691148e) [commit2](https://github.com/T-H-A-N-N/ebiznes25/commit/06e2bb6748587e8a15f82d97857be6ea4adc15e8)

3.5 rejestracja przez aplikację serwerową (bez Oauth2)

4.0 logowanie via Google OAuth2

4.5 logowanie via Facebook lub Github OAuth2

5.0 zapisywanie danych logowania OAuth2 po stronie serwera

Zadanie 9 ChatGPT bot

✅ 3.0 należy stworzyć po stronie serwerowej osobny serwis do łącznia z
chatGPT do usługi chat [commit](https://github.com/T-H-A-N-N/ebiznes25/commit/936597c6abfa1acc72afb0304dd60417fa2aad23)

3.5 należy stworzyć interfejs frontowy dla użytkownika, który
komunikuje się z serwisem; odpowiedzi powinny być wysyałen do
frontendowego interfejsu

4.0 stworzyć listę 5 różnych otwarć oraz zamknięć rozmowy

4.5 filtrowanie po zagadnieniach związanych ze sklepem (np.
ograniczenie się jedynie do ubrań oraz samego sklepu) do GPT

5.0 filtrowanie odpowiedzi po sentymencie
