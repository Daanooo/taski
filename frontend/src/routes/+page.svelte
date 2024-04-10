<script>
    import {
        NewTask,
        GetById,
    } from "../lib/wailsjs/go/controller/TaskController";
</script>

<button on:click={NewTask}>Something</button>

<div>
    {#await GetById(1)}
        <div>loading</div>
    {:then data}
        <div>{data.Description}</div>
    {/await}
</div>
