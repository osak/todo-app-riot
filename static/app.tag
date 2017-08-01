<app>
    <div each={this.todos}>
        <todo title={title} completed={completed}></todo>
    </div>
    <script>
        setTimeout(() => fetch("http://localhost:8080/todos")
            .then((res) => {
                return res.json();
            }).then((json) => {
                this.todos = json;
                this.update();
            }), 1000);
    </script>
</app>