<script>
    import Task from "../lib/components/Task.svelte";
    import {
        NewTask,
        GetAll,
    } from "../lib/wailsjs/go/controller/TaskController";

    const newTask = () => {
        console.log(NewTask("wtf"));
    };
</script>

<button on:click={newTask}>Something</button>

<div>
    {#await GetAll()}
        <div>loading tasks</div>
    {:then tasks}
        {#each tasks as task}
            <Task
                description={task.Description}
                completed={task.Completed}
                id={task.ID}
            />
        {/each}
    {/await}
</div>
