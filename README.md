# Алгоритм PAM

Рассмотрена и реализована усовершенствованная версия одного из распространенных алгоритмов кластеризации графов – PAM (Partitioning Around Medoids). 
Разбиение на кластеры зависит от входных данных, в отличие от базового метода, где количество кластеров считается изначально известным.

## Формат входных и выходных данных
Входные данные будут содержаться в файле in.txt, выходные – в out.txt.

### Формат входных данных:
В первой строке записано число k_max (int) – верхняя граница кластеризации..
Затем в формате yaml сам граф в виде списков смежности

### Формат выходных данных:
Кластеры списками смежностей в формате yaml (k-шт.)

смотреть в test_data

### Инструкция пользования
В зависимости от платформы скачать необходимый файл из папки executable_files. 
В консоли вводить: `./your_exefile input_file.txt output_file.txt`
Для запуска тестов в папке с test_coverage.go вводит: `go test -v`

### Пример

Входной файл    | Выходной файл
----------------|---------------
6               |1:
- 1:            |- 1
    3: 2        |- 2
    5: 15       |- 3
- 2:            |2:
    3: 1        |- 4
- 3:            |- 5
    1: 2        |- 6
    2: 1        |
- 4:            |
    5: 1        |
- 5:            |
    1: 15       |
    6: 4        |
    4: 1        |
- 6:            |
    5: 4        |
