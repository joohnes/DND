<script lang="ts">
    import { onMount } from 'svelte';
  	import { HOST } from "$lib/host";
	import Item from './Item.svelte';

	var eq: any
	onMount(async function () {
		loadData()
	});

	async function loadData() {
		const res = await fetch(HOST + "player/items/" + id)
		const data = await res.json()
		eq = data
	}

	let restartKey = {};
	const restart = () => {
		restartKey = {}
		loadData()
	}

	export let id: number = 0;
</script>
{#key restartKey}
<div class="flex flex-col justify-center items-center max-w-none">
	{#if eq != undefined}
	{#if eq.equipped != undefined && eq.equipped != null && eq.equipped.length != 0}
	<span class="text-xl">Equipped</span>
	<div class="flex flex-row flex-wrap gap-4 my-4">
			{#each eq.equipped as equipped}
				<Item id={equipped} playerID={id} playerView={true} eqView={true} equipped={true} restart={restart}/>
			{/each}
	</div>
	{/if}
	{#if eq.items != undefined && eq.items != null && eq.items.length != 0}
	<span class="text-xl">Inventory</span>
	<div class="flex flex-row flex-wrap gap-4 my-4">
			{#each eq.items as item}
			<div>
				<Item id={item} playerID={id} playerView={true} eqView={true} equipped={false} restart={restart}/>
			</div>
			{/each}
	</div>
	{/if}
	{/if}
</div>
{/key}