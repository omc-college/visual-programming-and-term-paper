# Resource management

Админка (Admin panel, back office)

### Пользовательская история (Use Case) group creation

    Как администратор, я хочу иметь возможность добавлять учебную группу.
    AC:
        При нажатии кнопки "добавить группу", откроется форма с следующими полями:
         * Group Name: required / max 255
         * Specialization
         * Year of education
         * Group number / (example: SD-31 or SD-32)
         * Curator
         * Students / optional(you can skip this step and add students later)
        При нажатии на кнопку "Сохоранить"
            - установка пароля (переиспользуем функциональность reset password)
