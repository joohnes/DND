<script>
    import Menu from "../components/menu.svelte";
</script>
<div style="display:flex">
    <svelte:component this={Menu}/>
    <div style="position:relative;left:7rem;">
        <slot />
    </div>
</div>

<style>
    @import url('https://fonts.googleapis.com/css2?family=Poppins:wght@300;400;500;600;700&display=swap');
    * {
        font-family: Poppins;
        font-size: 0.8rem;
    }
</style>