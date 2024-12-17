<script lang="ts">
	import ItemView from './Item.svelte';

	let restartKey = {};
	const restart = () => {
		restartKey = {}
	}

	let slotMap = new Map<number, string>([
		[1, "Head"],
		[2, "Chest"],
		[3, "Shoulders"],
		[4, "Hands"],
		[5, "Legs"],
		[6, "Feet"],
	])

	export let id: number = 0;
	export let items: Array<Item> | undefined;
	export let equipped: Array<Item> | undefined;
</script>
{#key restartKey}
	<div class="flex flex-col justify-center items-center max-w-none">
		{#if equipped != undefined && equipped.length > 0}
			<span class="text-2xl">Equipped</span>
			{#each equipped as eq}
				<div class="flex flex-row flex-wrap gap-4 my-4">
					<div class="flex flex-col">
						<span class="text-xl text-center mb-2">
							{#if eq.slot != undefined}
								{slotMap.get(eq.slot)}
							{/if}
						</span>
						<ItemView item={eq} playerID={id} playerView={true} eqView={true} equipped={true} restart={restart}/>
					</div>
				</div>
			{/each}
		{/if}
		{#if items != undefined && items.length > 0}
			<span class="text-2xl">Inventory</span>
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