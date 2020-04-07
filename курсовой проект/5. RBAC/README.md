# Role Based Access Control

### модуль RBAC





### Пользовательская история (Use Case)
As an admin I should be able:
    * Create new Roles based on Role Template
    * Update Roles
    * Delete Roles

### Техническая история (Technical Story)
Сущность Role -> имя фичи -> списку эндпоинтов
```golang
type Role struct{
    entries []FeatureEntry    
}

type FeatureEntry struct {
    featureName string
    endpoints []Endpoint
    accessGranted bool
}

```










###План работ
Любая функиональность это просто набор HTTP хэндлеров. 
    - Фича редактирование пользователя -> `PUT /users/{id}`
1. собрать список всех фич со всех микросервисов и создать Role Template
2. создать rbac сервис CRUD сервис для сущности Role, RoleTemplate
3. http middleware на golang парсить HTTP запросы и принимать решение об аторизации
4. модуль который при получении адейта ролей будет обновлять внутренний авторизационный кеш 
