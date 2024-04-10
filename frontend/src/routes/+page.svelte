<script>
    import {
        NewTask,
        GetById,
    } from "../lib/wailsjs/go/controller/TaskController";

    function newTask() {
        console.log(NewTask("test"));
        console.log(GetById(1));
    }
</script>

<h1>Welcome to SvelteKit</h1>
<p>
    Visit <a href="https://kit.svelte.dev">kit.svelte.dev</a> to read the documentation
</p>

<button on:click={newTask}>Something</button>
