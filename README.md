# Алгоритм PAM
Рассмотрена и реализована усовершенствованная версия одного из распространенных алгоритмов кластеризации графов – PAM (Partitioning Around Medoids). 
Разбиение на кластеры зависит от входных данных, в отличие от базового метода, где количество кластеров считается изначально известным.

## Формат входных и выходных данных
Входные данные будут содержаться в файле in.txt, выходные – в out.txt.

### Формат входных данных:
В первой строке записано число k_max (int) – верхняя граница кластеризации.
Затем в формате yaml сам граф в виде списков смежности

### Формат выходных данных:
Кластеры списками смежностей в формате yaml (k-шт.)

Примеры смотреть в test_data

### Инструкция пользования
Для адекватной работы программы необходима версия языка Go 1.8 и старше, 
иначе возникает ошибка, связанная с отсутствием sort.Slice в предыдущих версиях (рекомендуется установить последнюю).
Посмотреть версию – команда: `go version`
Для того чтобы собрать проект, пишем:
`go build main.go pam.go floyd_warshall.go init.go`
(порядок модулей не важен)
Если не видит пакет `gopkg.in/yaml.v2`, устанавливаем его:
`go get –u gopkg.in/yaml.v2`
Далее программа готова к использованию. Пишем:
`«executable_file» «inputFileName» «outputFileName»`
Для запуска тестов в одной папке с test_coverage.go выполнить команду:
`go test –v`
