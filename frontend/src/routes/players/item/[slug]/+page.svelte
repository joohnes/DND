<script lang="ts">
    import Item from "../../../../components/Item.svelte";
    /** @type {import('./$types').PageData} */
    export let data;
    const id = data.data
</script>

<Item id={id} playerView={true}/>