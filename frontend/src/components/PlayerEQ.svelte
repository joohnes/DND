<script lang="ts">
	import ItemView from './Item.svelte';

	let restartKey = {};
	const restart = () => {
		restartKey = {}
	}

	export let id: number = 0;
	export let items: Array<Item> | undefined;
	export let equipped: Array<Item> | undefined;
</script>
{#key restartKey}
<div class="flex flex-col justify-center items-center max-w-none">
	{#if equipped != undefined && equipped.length > 0}
	<span class="text-xl">Equipped</span>
	<div class="flex flex-row flex-wrap gap-4 my-4">
			{#each equipped as eq}
				<ItemView item={eq} playerID={id} playerView={true} eqView={true} equipped={true} restart={restart}/>
			{/each}
	</div>
	{/if}
	{#if items != undefined && items.length > 0}
	<span class="text-xl">Inventory</span>
	<div class="flex flex-row flex-wrap gap-4 my-4">
			{#each items as item}
			<div>
				<ItemView item={item} playerID={id} playerView={true} eqView={true} equipped={false} restart={restart}/>
			</div>
			{/each}
	</div>
	{/if}
</div>
{/key}