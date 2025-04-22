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
