<script>
    export let id;
    export let description;
    export let completed;
</script>

<div class="flex flex-row h-12">
    <div class="basis-1/4">
        <p>{id}</p>
    </div>
    <div class="basis-1/2">
        <p>{description}</p>
    </div>
    <div class="basis-1/4">
        <p>{completed}</p>
    </div>
</div>
