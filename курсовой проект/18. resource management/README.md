# Resource management

Админка (Admin panel, back office)


### Пользовательская история (Use Case) group creation
    Как администратор, я хочу иметь возможность добавлять учебную группу.
    AC:
        При нажатии кнопки "добавить группу", откроется форма с следующими полями:
         * Group Name: required / max 255
         * Specialization
         * Year of education
        При нажатии на кнопку "Сохоранить"
            - установка пароля (переиспользуем функциональность reset password)
 
