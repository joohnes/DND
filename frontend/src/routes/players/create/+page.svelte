<script lang="ts">
  import { HOST } from "$lib/host";
  import { onMount, onDestroy } from 'svelte';

  var menuItem: HTMLElement
  onMount(()=> {
    menuItem = document.getElementById("menu-players-create")!
		menuItem.classList.remove("btn-outline")

  })

  onDestroy(() => {
		menuItem.classList.add("btn-outline")
	})

  const sendData = (e: any) => {
      let p: Player
      const formData = new FormData(e.target)
      p = {
        name: formData.get("name")!.toString(),
        level: parseInt(formData.get("level")!.toString()),
        class: formData.get("class")!.toString(),
        race: formData.get("race")!.toString(),
        subrace: formData.get("subrace")!.toString(),
        strength: parseInt(formData.get("strength")!.toString()),
        endurance: parseInt(formData.get("endurance")!.toString()),
        perception: parseInt(formData.get("perception")!.toString()),
        intelligence: parseInt(formData.get("intelligence")!.toString()),
        agility: parseInt(formData.get("agility")!.toString()),
        accuracy: parseInt(formData.get("accuracy")!.toString()),
        charisma: parseInt(formData.get("charisma")!.toString())
      }

      fetch(HOST + "player", {
        method: "POST",
        body: JSON.stringify(p)
      })
      window.location.href = window.location.origin + "/players"
    }

</script>
<div class="h-screen my-auto items-center py-1 px-8 m-2 flex flex-col drop-shadow-xl">
<form on:submit|preventDefault={sendData}>
    <fieldset class="border border-slate-900 rounded-lg p-5 shadow-2xl">
      <div class="flex gap-20">
        <div>
          <label  for="name">Name</label>  
          <div>
           <input id="name" name="name" type="text" class="input input-bordered w-full max-w-md">
          </div>
    
          <label class="" for="level">Level</label>  
          <div>
            <input id="level" name="level" type="number" class="input input-bordered w-full max-w-xs">
          </div>

          <label for="class">Class</label>  
          <div>
            <select name="class" class="select select-bordered w-full max-w-xs">
              <option disabled selected>Choose class</option>
              <option>Common</option>
              <option>Rare</option>
              <option>Epic</option>
              <option>Legendary</option>
              <option>Artefact</option>
            </select>
          </div>
        
          <label for="race">Race</label>  
          <div>
            <select name="race" class="select select-bordered w-full max-w-xs">
              <option disabled selected>Choose race</option>
              <option>Common</option>
              <option>Rare</option>
              <option>Epic</option>
              <option>Legendary</option>
              <option>Artefact</option>
            </select>
          </div>
        
        
          <label for="subrace">Subrace</label>  
          <div>
            <select name="subrace" class="select select-bordered w-full max-w-xs">
              <option disabled selected>Choose subrace</option>
              <option>Common</option>
              <option>Rare</option>
              <option>Epic</option>
              <option>Legendary</option>
              <option>Artefact</option>
            </select>
          </div>

        </div>
        <div>
          <label  for="strength">Strength</label>  
          <div >
            <input id="strength" name="strength" type="number" class="input input-bordered w-full max-w-xs">
          </div>
        
        
          <label  for="endurance">Endurance</label>  
          <div >
            <input id="endurance" name="endurance" type="number" class="input input-bordered w-full max-w-xs">
          </div>
        
        
          <label  for="perception">Perception</label>  
          <div >
            <input id="perception" name="perception" type="number" class="input input-bordered w-full max-w-xs">
          </div>
        
        
          <label  for="intelligence">Intelligence</label>  
          <div >
          <input id="intelligence" name="intelligence" type="number" placeholder="" class="input input-bordered w-full max-w-xs">
          </div>
        
        
          <label  for="agility">Agility</label>  
          <div >
          <input id="agility" name="agility" type="number" placeholder="" class="input input-bordered w-full max-w-xs">
          </div>
        
        
          <label  for="accuracy">Accuracy</label>  
          <div >
          <input id="accuracy" name="accuracy" type="number" placeholder="" class="input input-bordered w-full max-w-xs">
          </div>
        
        
          <label  for="charisma">Charisma</label>  
          <div >
          <input id="charisma" name="charisma" type="number" placeholder="" class="input input-bordered w-full max-w-xs">
          </div>
        </div>
      </div>
      <div class="flex justify-center mt-4">
        <button class="btn btn-outline btn-primary px-8"><input type="submit" value="Create"></button>
      </div>
    </fieldset>
    </form>
  </div>
    