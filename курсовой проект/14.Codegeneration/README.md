# Codegeneration

### модуль кодогенерации структур моделей, роутера, валидаторов
Задача заключается в том, тчобы генерировать те части веб преложения которые фактически реализутся на и полностью описаны на основании openAPI спецификации веб сервиса.




### Техническая история (Technical Story)



```yaml
      Endpoint:
         type: object
         properties:
            id:
               type: string
               maxLength: 255
               description: Endpoint id
            path:
               type: string
               maxLength: 255
               description: Endpoint path
            method:
               type: string
               maxLength: 255
               description: Endpoint method
```

```go
type Endpoint struct {
    ID string
    Path string
    Method string
}
```






###План работ
1. Написать тестовый веб сервис (ссылка на github repository)
2. Рефакторинг, чтобы выделить основные компоненты (маппинг OpenAPI спеки на сущности веб сервера)
3. Распарсить OpenAPI спецификацию.
---------
4. начать с генерации можелей используя этот пакет - https://golang.org/pkg/text/template/
