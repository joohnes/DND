<script lang="ts">
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
      }

      fetch(HOST + "/item", {
        method: "POST",
        body: JSON.stringify(i),
      })
      window.location.href = window.location.origin + "/items"
    }
</script>

<form class="form-horizontal" on:submit|preventDefault={sendData}>
    <fieldset>
    <div class="form-group">
      <label class="col-md-4 control-label" for="name">Name</label>  
      <div class="col-md-4">
      <input id="name" name="name" type="text" class="form-control input-md">
      </div>
    </div>
    
    <div class="form-group">
      <label class="col-md-4 control-label" for="description">Description</label>  
      <div class="col-md-4">
      <input id="description" name="description" type="text" class="form-control input-md">
      </div>
    </div>
    
    <div class="form-group">
      <label class="col-md-4 control-label" for="ability">Ability</label>  
      <div class="col-md-4">
      <input id="ability" name="ability" type="text" class="form-control input-md">
      </div>
    </div>

    <div class="form-group">
      <label class="col-md-4 control-label" for="rarity">Rarity</label>  
      <div class="col-md-4">
      <select name="rarity">
        <option>Common</option>
        <option>Rare</option>
        <option>Legendary</option>
        <option>Artefact</option>
      </select>
      </div>
    </div>
    
    <div class="form-group">
      <label class="col-md-4 control-label" for="strength">Strength</label>  
      <div class="col-md-4">
      <input id="strength" name="strength" type="number" placeholder="" class="form-control input-md">
      </div>
    </div>
    
    <div class="form-group">
      <label class="col-md-4 control-label" for="endurance">Endurance</label>  
      <div class="col-md-4">
      <input id="endurance" name="endurance" type="number" placeholder="" class="form-control input-md">
      </div>
    </div>
    
    <div class="form-group">
      <label class="col-md-4 control-label" for="perception">Perception</label>  
      <div class="col-md-4">
      <input id="perception" name="perception" type="number" placeholder="" class="form-control input-md">
      </div>
    </div>
    
    <div class="form-group">
      <label class="col-md-4 control-label" for="intelligence">Intelligence</label>  
      <div class="col-md-4">
      <input id="intelligence" name="intelligence" type="number" placeholder="" class="form-control input-md">
      </div>
    </div>
    
    <div class="form-group">
      <label class="col-md-4 control-label" for="agility">Agility</label>  
      <div class="col-md-4">
      <input id="agility" name="agility" type="number" placeholder="" class="form-control input-md">
      </div>
    </div>
    
    <div class="form-group">
      <label class="col-md-4 control-label" for="accuracy">Accuracy</label>  
      <div class="col-md-4">
      <input id="accuracy" name="accuracy" type="number" placeholder="" class="form-control input-md">
      </div>
    </div>
    
    <div class="form-group">
      <label class="col-md-4 control-label" for="charisma">Charisma</label>  
      <div class="col-md-4">
      <input id="charisma" name="charisma" type="number" placeholder="" class="form-control input-md">
      </div>
    </div>
    <div class="form-group">
      <label class="col-md-4 control-label" for="quantity">Quantity</label>  
      <div class="col-md-4">
        <input id="quantity" name="quantity" type="number" placeholder="" class="form-control input-md">
      </div>
    </div>
    <input type="submit" value="Create">
    </fieldset>
    </form>
    