var todos = [];

window.onload = function () {
    GetAll();

    $('[data-toggle="tooltip"]').tooltip();

    function formatDate(date) {
        let day = date.getDate();
        let month = date.getMonth();
        let year = date.getFullYear();

        // Get ordinal suffix for the day
        let ordinal;
        if (day % 10 == 1 && day != 11) {
            ordinal = "st";
        } else if (day % 10 == 2 && day != 12) {
            ordinal = "nd";
        } else if (day % 10 == 3 && day != 13) {
            ordinal = "rd";
        } else {
            ordinal = "th";
        }

        // Array of month names for easier retrieval
        let monthNames = ["Jan", "Feb", "Mar", "Apr", "May", "Jun", "Jul", "Aug", "Sep", "Oct", "Nov", "Dec"];

        return `${day}${ordinal} ${monthNames[month]} ${year}`;
    }


    var currentDate = formatDate(new Date());

    $(".due-date-button").datepicker({
        format: "dd/mm/yyyy",
        autoclose: true,
        todayHighlight: true,
        startDate: currentDate,
        orientation: "bottom right"
    });

    $(".due-date-button").on("click", function (event) {
        $(".due-date-button")
            .datepicker("show")
            .on("changeDate", function (dateChangeEvent) {
                $(".due-date-button").datepicker("hide");
                $(".due-date-label").text(formatDate(dateChangeEvent.date));
            });
    });
    
    let addButton = document.querySelector('.btn.btn-primary');
    addButton.addEventListener('click', function() {
        let input = document.querySelector('.add-todo-input')
        let inputValue = input.value;
        if (inputValue === "") {
            alertify.warning('Please enter todo title')
            return;
        }
        input.value = "";
        let dueDate = document.querySelector('.due-date-label').textContent;
        $(".due-date-label").text("");
        let todo = {
            id: 0,
            title: inputValue,
            is_completed: false,
            created_at: formatDate(new Date()),
            due_date: dueDate
        }
        CreateTodo(todo);
        alertify.success('todo created');
    });
};

function CreateTodo(todo = {
    id: 0,
    title: "",
    is_completed: false,
    created_at: "",
    due_date: "",
}) {
    fetch('/api/todos/create', {
        method: 'POST',
        headers: {
            'Content-Type': 'application/json'
        },
        body: JSON.stringify(todo)
    })
        .then(response => GetAll())
}

function GetAll() {
    fetch('/api/todos/all', {
        method: 'GET',
        headers: {
            'Content-Type': 'application/json'
        },
    })
        .then(response => response.json())
        .then(data => {
            todos = data;
            generateTodoList()
        } )
        .catch(error => console.error(error));
}

function Delete(id) {
    fetch(`/api/todos/delete/${id}`, {
        method: 'POST',
        headers: {
            'Content-Type': 'application/json'
        },
    })
        .then(response => GetAll())
        .then(data => {
            console.log(data);
        } )
        .catch(error => console.error(error));
}

function Update(todo = {
    id: 0,
    title: "",
    is_completed: false,
    created_at: "",
    due_date: "",
}) {
    fetch(`/api/todos/update`, {
        method: 'POST',
        headers: {
            'Content-Type': 'application/json'
        },
        body: JSON.stringify(todo)
    })
        .then(response => response.json())
        .then(data => {
            GetAll();
        } )
        .catch(error => console.error(error));
}

function Complete(id) {
    fetch(`/api/todos/complete/${id}`, {
        method: 'POST',
        headers: {
            'Content-Type': 'application/json'
        },
    })
        .then(response => GetAll())
        .then(data => {
            console.log(data);
        } )
        .catch(error => console.error(error));
}

function generateTodoList() {
    var todoContainer = document.querySelector(".todos");
    todoContainer.innerHTML = "";

    todos.forEach(function(todo) {
        var todoItem = document.createElement("div");
        todoItem.className = "row px-3 align-items-center todo-item rounded";

        // Create HTML structure for todo item using the todo properties
        todoItem.innerHTML = `
            <div class="col-auto m-1 p-0 d-flex align-items-center editing" id="todo-id-${todo.id}">
                <h2 class="m-0 p-0">
                    <i class="fa ${todo.is_completed ? 'fa-check-square-o' : 'fa-square-o'} text-primary btn m-0 p-0 check-me" data-toggle="tooltip" data-placement="bottom" data-id="${todo.id}" data-checked="${todo.is_completed}" title="${todo.is_completed ? 'Mark as todo' : 'Mark as complete'}"></i>
                </h2>
            </div>
            <div class="col px-1 m-1 d-flex align-items-center">
                <input type="text" class="form-control form-control-lg border-0 edit-todo-input rounded px-3" value="${todo.title}" title="${todo.title}" data-id="${todo.id}"/>
            </div>
            <div class="col-auto m-1 p-0 px-3 ${todo.due_date ? '' : 'd-none'}">
                <div class="row">
                    <div class="col-auto d-flex align-items-center rounded bg-white border border-warning">
                        <i class="fa fa-hourglass-2 my-2 px-2 text-warning btn" data-toggle="tooltip" data-placement="bottom" title="Due on date"></i>
                        <h6 class="text my-2 pr-2">${todo.due_date}</h6>
                    </div>
                </div>
            </div>
            <div class="col-auto m-1 p-0 todo-actions">
                <div class="row d-flex align-items-center justify-content-end">
                    <h5 class="m-0 p-0 px-2">
                        <i class="fa fa-trash-o text-danger btn m-0 p-0" data-toggle="tooltip" data-placement="bottom" title="Delete todo" data-id="${todo.id}"></i>
                    </h5>
                </div>
                <div class="row todo-created-info">
                    <div class="col-auto d-flex align-items-center pr-2">
                        <i class="fa fa-info-circle my-2 px-2 text-black-50 btn" data-toggle="tooltip" data-placement="bottom" title="Created date"></i>
                        <label class="date-label my-2 text-black-50">${todo.created_at}</label>
                    </div>
                </div>
            </div>
        `;

        todoContainer.appendChild(todoItem);
    });
    const deleteBtns = document.querySelectorAll(".fa-trash-o");
    deleteBtns.forEach(btn => {
        btn.addEventListener("click", function() {
            alertify.confirm("Удаление", "Вы уверены, что хотите удалить?",
                function(){
                    const elementId = btn.getAttribute("data-id");
                    Delete(elementId);
                    alertify.success('Deleted');
                },
                function(){
                    alertify.error('Cancel');
                });
        });
    });
    
    const completeBtns = document.querySelectorAll(".check-me");
    completeBtns.forEach(btn => {
        btn.addEventListener("click", function() {
            const elementId = this.getAttribute("data-id");
            Complete(elementId);
            alertify.success('todo completed');
        });
    });
    
    const editIcons = document.querySelectorAll('.fa-pencil');

    editIcons.forEach(icon => {
        icon.addEventListener('click', function() {
            const todoElement = icon.closest('.todo-item');
            console.log(todoElement);
            todoElement.classList.add('editing');
        });
    });
    
    // edit-todo-input
    const editInputs = document.querySelectorAll('.edit-todo-input');
    editInputs.forEach(input => {
        input.addEventListener('keyup', function(event) {
            const todoId = input.getAttribute('data-id');
            const todoTitle = input.value;
            const todo = {
                id: parseInt(todoId),
                title: todoTitle,
                is_completed: false,
                created_at: "",
                due_date: "",
            }
            Update(todo);
            alertify.success('todo updated');
        });
    })
}
