<script lang="ts">
  import { HOST } from "$lib/host";
	import { onMount, onDestroy } from 'svelte';

	var menuItem: HTMLElement
  onMount(()=> {
    menuItem = document.getElementById("menu-items-create")!
		menuItem.classList.remove("btn-outline")

  })

  onDestroy(() => {
		menuItem.classList.add("btn-outline")
	})

  let slotMap = new Map<string, number>([
		["Head", 1],
		["Chest", 2],
		["Shoulders", 3],
		["Hands", 4],
		["Legs", 5],
		["Feet", 6],
	])

  const sendData = (e: any) => {
      let i: Item
      const formData = new FormData(e.target)
      i = {
        name: formData.get("name")!.toString(),
        description: formData.get("description")!.toString(),
        ability: formData.get("ability")!.toString(),
        rarity: formData.get("rarity")!.toString(),
        strength: parseInt(formData.get("strength")!.toString()),
        endurance: parseInt(formData.get("endurance")!.toString()),
        perception: parseInt(formData.get("perception")!.toString()),
        intelligence: parseInt(formData.get("intelligence")!.toString()),
        agility: parseInt(formData.get("agility")!.toString()),
        accuracy: parseInt(formData.get("accuracy")!.toString()),
        charisma: parseInt(formData.get("charisma")!.toString()),
        quantity: parseInt(formData.get("quantity")!.toString()),
        attack: parseInt(formData.get("attack")!.toString()),
        defense: parseInt(formData.get("defense")!.toString()),
        slot: slotMap.get(formData.get("slot")!.toString())
      }

      fetch(HOST + "item", {
        method: "POST",
        body: JSON.stringify(i),
      })
      window.location.href = window.location.origin + "/items"
    }
</script>

<div class="h-screen my-auto items-center py-1 px-8 m-2 flex flex-col drop-shadow-xl">
<form on:submit|preventDefault={sendData}>
    <fieldset class="border border-slate-900 rounded-lg p-5 shadow-2xl">
    <div class="flex gap-20">
      <div>
        <label for="name">Name</label>  
        <div>
        <input id="name" name="name" type="text" class="input input-bordered w-full max-w-xs">
        </div>
      
        <label for="description">Description</label>  
        <div>
        <input id="description" name="description" type="text" class="input input-bordered w-full max-w-xs">
        </div>
      
        <label for="ability">Ability</label>  
        <div>
        <input id="ability" name="ability" type="text" class="input input-bordered w-full max-w-xs">
        </div>

        <label for="attack">Attack</label>  
        <div>
        <input id="attack" name="attack" type="number" class="input input-bordered w-full max-w-xs">
        </div>

        <label for="defense">Defense</label>  
        <div >
        <input id="defense" name="defense" type="number" class="input input-bordered w-full max-w-xs">
        </div>

        <label for="rarity">Rarity</label>  
        <div>
          <select name="rarity" class="select select-bordered w-full max-w-xs">
            <option disabled selected>Choose Rarity</option>
            <option>Common</option>
            <option>Rare</option>
            <option>Epic</option>
            <option>Legendary</option>
            <option>Artefact</option>
          </select>
        </div>

        <label for="slot">Slot</label>  
        <div>
          <select name="slot" class="select select-bordered w-full max-w-xs">
            <option disabled selected>Choose slot</option>
            <option>Head</option>
            <option>Chest</option>
            <option>Shoulders</option>
            <option>Hands</option>
            <option>Legs</option>
            <option>Feet</option>
          </select>
        </div>
      </div>
      
      <div>
        <label for="strength">Strength</label>  
        <div >
        <input id="strength" name="strength" type="number" placeholder="" class="input input-bordered w-full max-w-xs">
        </div>
      
        <label for="endurance">Endurance</label>  
        <div >
        <input id="endurance" name="endurance" type="number" placeholder="" class="input input-bordered w-full max-w-xs">
        </div>
      
        <label for="perception">Perception</label>  
        <div >
        <input id="perception" name="perception" type="number" placeholder="" class="input input-bordered w-full max-w-xs">
        </div>
      
        <label for="intelligence">Intelligence</label>  
        <div >
        <input id="intelligence" name="intelligence" type="number" placeholder="" class="input input-bordered w-full max-w-xs">
        </div>
      
        <label for="agility">Agility</label>  
        <div >
        <input id="agility" name="agility" type="number" placeholder="" class="input input-bordered w-full max-w-xs">
        </div>
      
        <label for="accuracy">Accuracy</label>  
        <div >
        <input id="accuracy" name="accuracy" type="number" placeholder="" class="input input-bordered w-full max-w-xs">
        </div>
      
        <label for="charisma">Charisma</label>  
        <div >
        <input id="charisma" name="charisma" type="number" placeholder="" class="input input-bordered w-full max-w-xs">
        </div>
        
        <label for="quantity">Quantity</label>  
        <div >
          <input id="quantity" name="quantity" type="number" placeholder="" class="input input-bordered w-full max-w-xs">
        </div>
      </div>
      </div>
      <button class="btn"><input type="submit" value="Create"></button>
    </fieldset>
    </form>
    </div>