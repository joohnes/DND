<script lang="ts">
	import { onMount } from 'svelte';
	import Modal from './Modal.svelte'
	import PlayerEQ from './PlayerEQ.svelte';
  	import { HOST } from "$lib/host";

	export let p: PlayerResponse;
	let min = 1
	onMount(async function () {
		if (GetAlcoRange().length > 5) {
			min = 0
		}
	});

	async function UpdateHPMANA() {
		let data = {
			hp: p.health,
			mana: p.mana
		}
		fetch(HOST + "player/hpmana/" + p.id, {
        method: "PUT",
        body: JSON.stringify(data)
      })
	  showModal = false
      restart()
	}

	async function DeletePlayer() {
		fetch(HOST + "player/" + p.id, {
        method: "DELETE",
		})
		restart()
	}


	function GetAlcoRange() {
		if (p.alcohol_level == undefined) {
			return [1, 2, 3, 4, 5]
		}
    const minLevel = 5; // Minimum value for each step
    const maxLevel = 200; // Maximum value for each step

    if (p.alcohol_level < minLevel || p.alcohol_level > maxLevel) {
      return [1, 2, 3, 4, 5]; // Default range if out of bounds
    }

    const step = Math.floor((p.alcohol_level - minLevel) / 5);
    const start = minLevel + step * 5;

    return Array.from({ length: 5 }, (_, i) => start + i);
  }

	function GetAlcoColor() {
		if (p.alcohol_level == undefined || p.alcohol_level <= 1) {
			return ""
		} else if (p.alcohol_level == 2) {
			return "range-success"
		} else if (p.alcohol_level < 10) {
			return "range-warning"
		} else if (p.alcohol_level == 10) {
			return "range-info"
		} else {
			return "range-error"
		}
	}

	function ConvertAlcoLevel() {
		if (p.alcohol_level == undefined) {
			return 0
		}
		let lvl = p.alcohol_level
		while (true) {
			if (lvl > 5) {
				lvl -= 5
			} else {
				if (lvl == 5) {
					return lvl - 5
				}
				return lvl + 1
			}
		}
	}

	async function Zgon() {
		await fetch(HOST + "player/zgon/" + p.id, {
			method: "POST"
		})
		restart()
	}

	let showModal: boolean;
	let modalUpdate: boolean;
	let modalEq: boolean;
	export let restart: Function;
</script>
{#if modalUpdate}
<Modal bind:showModal>
	<div class="flex flex-col justify-center items-center">
	<h2>Enter new HP and MANA</h2>
	<div class="flex flex-col py-2 gap-4">
		<div class="flex gap-4 justify-between">
			<label for="hp" class="my-auto">HP</label>
			<input type="number" id="hp" name="hp" class="input input-sm input-bordered w-20" bind:value={p.health}>
		</div>
		<div class="flex gap-4 justify-between">
			<label for="mana" class="my-auto">MANA</label>
			<input type="number" id="mana" name="mana" class="input input-sm input-bordered w-20" bind:value={p.mana}>
		</div>
	</div>
	<button class="btn btn-outline btn-success px-12 mb-2" on:click={()=> {UpdateHPMANA()}}>Update</button>
</div>
</Modal>
{:else if !modalUpdate && !modalEq}
<Modal bind:showModal>
	<div class="flex flex-col justify-center items-center">
	<h2>Delete player?</h2>
	<button class="btn" on:click={()=> {DeletePlayer()}}>DELETE</button>
</div>
</Modal>
{:else if modalEq}
<Modal bind:showModal fullScreen={true}>
	{#if (p.items != undefined && p.items.length > 0) || (p.equipped != undefined && p.equipped.length > 0)}
		<PlayerEQ items={p.items} equipped={p.equipped}/>
	{:else}
		<div class="text-lg">No Items in inventory</div>
	{/if}
</Modal>
{/if}
{#if p != undefined}
	<div class="player {p.zgon ? "grayscale" : ""}">
		<div class="info">
			<div>
				<!-- svelte-ignore a11y-click-events-have-key-events -->
				<!-- svelte-ignore a11y-no-static-element-interactions -->
				<span on:click={()=>{modalUpdate=false;modalEq=true;showModal = true}}>[EQ]</span>
				<!-- svelte-ignore a11y-click-events-have-key-events -->
				<!-- svelte-ignore a11y-no-static-element-interactions -->
				<span on:click={()=>{modalUpdate=true;modalEq=false;showModal = true}}>
					<span class="text-red-500">HP<span class="text-blue-600">MANA</span></span>
				</span>
				{#if p.is_holder}
					<a href={window.location.origin+"/bag"}>🎒</a>
				{/if}
			</div>
			<div>
				<!-- svelte-ignore a11y-no-static-element-interactions -->
				<!-- svelte-ignore a11y-click-events-have-key-events -->
				<span on:click={Zgon}>💀</span>
				<a href={window.location.origin+"/players/update/"+p.id}>⚙️</a>
				<!-- svelte-ignore a11y-click-events-have-key-events -->
				<!-- svelte-ignore a11y-no-static-element-interactions -->
				<!-- svelte-ignore a11y-missing-attribute -->
				<a on:click={()=>{modalUpdate=false;modalEq=false;showModal=true}}>❌</a>
			</div>
		</div>
		<div class="cellHolder">
			<div class="underline underline-offset-1 name">{p.name}</div>
			<div class="badge badge-ghost badge-outline">LVL {p.level}</div>
		</div>
		<div class="cellHolder">
			<div class="badge badge-primary badge-outline">{p.class}</div>
			<div class="badge badge-warning badge-outline">{p.race}</div>
			<div class="badge badge-accent badge-outline">{p.subrace}</div>
		</div>
		<div class="cellHolder">
			<div class="badge badge-outline badge-secondary">HP {p.health}</div>
			{#if p.zgon}
				<div class="badge badge-outline">zgon</div>
			{/if}
			<div class="badge badge-outline badge-info">MANA {p.mana}</div>
		</div>

		{#if p.alcohol_level != undefined}
		<div class="cellHolder flex-col">
			<input type="range" disabled min={min} max="5" value={ConvertAlcoLevel()} class="range range-xs mb-1 {GetAlcoColor()}" step="1" />
				<div class="w-full flex justify-between text-xs px-2">
					{#each GetAlcoRange() as i}
						<span>{i}</span>
					{/each}
			</div>
		</div>
		{/if}

		<div class="cellHolder flex-col">
			<div class="statsCell">
				<div>Strength:</div>
				<div>{p.strength_printable}</div>
			</div>
			<div class="statsCell">
				<div>Endurance:</div>
				<div>{p.endurance_printable}</div>
			</div>
			<div class="statsCell">
				<div>Perception:</div>
				<div>{p.perception_printable}</div>
			</div>
			<div class="statsCell">
				<div>Intelligence:</div>
				<div>{p.intelligence_printable}</div>
			</div>
			<div class="statsCell">
				<div>Agility:</div>
				<div>{p.agility_printable}</div>
			</div>
			<div class="statsCell">
				<div>Accuracy:</div>
				<div>{p.accuracy_printable}</div>
			</div>
			<div class="statsCell">
				<div>Charisma:</div>
				<div>{p.charisma_printable}</div>
			</div>
		</div>
	</div>
{/if}

<style>
	.player {
		display: flex;
		flex-direction: column;
		width: 20rem;
		justify-content: space-between;
		padding: 1rem;
		border: 3px solid;
		border-radius: 1rem;
		@apply border-slate-500;
	}
	.cellHolder {
		display: flex;
		padding-top: 0.5rem;
		padding-bottom: 0.5rem;
		border-bottom: 1px solid;
		@apply border-slate-600;
		justify-content: space-between;
	}
	.statsCell {
		display: flex;
		flex-direction: row;
		justify-content: space-between;
	}
	.info {
		display: flex;
		justify-content: space-between;
	}
	.info > div > span, .info > div > a {
		cursor: pointer;
	}
	.name {
		width: 9rem;
		overflow-wrap: break-word;
	}
</style>
