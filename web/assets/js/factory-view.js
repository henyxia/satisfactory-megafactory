async function modalFactoryDelete_Send(id) {
  const response = await fetch('/api/factory/'+id, {
    method: 'DELETE',
    headers: {
      'Content-Type': 'application/json'
    },
  });

  try {
    const result = await response.json();
    if (!response.ok) {
        console.warn(result.error);
    } else {
      window.location.pathname = '/';
    }
  } catch(e) {
    console.error(e);
  }
}

async function modalFactoryEdit_Send(id) {
  data = {
      "id": id,
      "name": getInputValue("modal-factory-new-name"),
      "io_north": getInputValueInt("modal-factory-new-io-north"),
      "io_east": getInputValueInt("modal-factory-new-io-east"),
      "io_south": getInputValueInt("modal-factory-new-io-south"),
      "io_west": getInputValueInt("modal-factory-new-io-west"),
  };

  const response = await fetch('/api/factory', {
      method: 'POST',
      headers: {
        'Content-Type': 'application/json'
      },
      body: JSON.stringify(data)
  });
    
  try {
    const result = await response.json();
    if (!response.ok) {
        console.warn(result.error);
    } else {
        window.location.reload();
    }
  } catch(e) {
    console.error(e);
  }
}

async function modalFloorNew_Send(factory_id) {
  data = {
      "name": getInputValue("modal-floor-new-name"),
      "level": getInputValueInt("modal-floor-new-level"),
      "status": getInputValue("modal-floor-new-status"),
      "factory_id": factory_id,
  };

  const response = await fetch('/api/factory/floor', {
      method: 'POST',
      headers: {
        'Content-Type': 'application/json'
      },
      body: JSON.stringify(data)
    });

  try {
    const result = await response.json();
    if (!response.ok) {
        console.warn(result.error);
    } else {
        window.location.reload();
    }
  } catch(e) {
    console.error(e);
  }
}

function factoryView_ioCellContent(io) {
  content = '';
  content += '<div class="tags has-addons"><span class="tag is-dark">'+io.Building.Building.Name;
  content += '</span><span class="tag is-success">'+io.Building.Recipe.Name+'</span></div>';

  if (io.Building.Recipe.Output[0] != null) {
    throughput = 60;
    throughput *= io.Building.Recipe.Output[0].Qt;
    throughput /= io.Building.Recipe.Duration;

    content += '<div class="tags has-addons">';
    content += '<span class="tag is-dark">Speed</span>';
    content += '<span class="tag is-info">'+throughput+'/m</span>';
    content += '</div>';
  }

  if (io.Building.note != '') {
    content += '<div class="tags has-addons">';
    content += '<span class="tag is-dark">Note</span>';
    content += '<span class="tag is-primary">'+io.Building.note+'</span>';
    content += '</div>';
  }

  return content;
}

async function factoryView_LoadIOs(factory_id, floor_id) {
  const response = await fetch('/api/floor/'+floor_id+'/io', {
    method: 'GET',
    headers: {
      'Content-Type': 'application/json'
    },
  });

  try {
    const result = await response.json();
    if (!response.ok) {
        console.warn(result.error);
    } else {
      for (const io of result.ios) {
        ioID = 'io-'+io.direction+'-'+io.position+'-'+io.sub_position;
        document.getElementById(ioID).innerHTML = factoryView_ioCellContent(io);
      }

      for (const tr of document.getElementById('io-table').children[1].children) {
        for (const td of tr.children) {
          if (td.innerHTML != '') {
            continue;
          }

          cell = '<div class="modal-new-io-button has-text-centered btn" onclick="factoryView_NewIOShowModal(this);">';
          cell += "New IO";
          cell += '</div>';

          td.innerHTML = cell;
        }
      }

      document.getElementById('modal-io-new-floor-id').value = floor_id;
    }
  } catch(e) {
    console.error(e);
  }
}


async function factoryView_NewIOShowModal(caller) {
  document.getElementById('modal-io-new-position-coded').value = caller.parentNode.id;

  showModal('io-new');
}

async function factoryView_newIOGetBuildingRecipes() {
  building_id = document.getElementById('modal-io-new-building').value;
  if (building_id == '') {
    return;
  }
  const response = await fetch('/api/building/'+building_id+'/recipe', {
    method: 'GET',
    headers: {
      'Content-Type': 'application/json'
    },
  });

  try {
    const result = await response.json();
    if (!response.ok) {
        console.warn(result.error);
    } else {
      cell = '<option value="" selected>Select a recipe</option>';
      for (const recipe of result.recipes) {
        cell += '<option value ="'+recipe.Slug+'">'+recipe.Name+'</option>';
      }
      document.getElementById('modal-io-new-recipe').innerHTML = cell;
    }
  } catch(e) {
    console.error(e);
  }
}

async function factoryView_NewIOSend() {
  /* ex: io-0-7-3
   *        | | |
   *        | | + ---- subposition : 0 lower right, 1 higher right, 2 lower left, 3 higher left
   *        | +------ position
   *        +-------- direction    : 0 north, 1 east, 2 south, 3 west
   */
  position_coded = getInputValue('modal-io-new-position-coded').split('-');

  floor_id = getInputValue('modal-io-new-floor-id');

  data = {
    // iobuilding
    "note": getInputValue("modal-io-new-note"),
    "building_id": getInputValueInt("modal-io-new-building"),
    "recipe_slug": getInputValue("modal-io-new-recipe"),

    // io
    "direction": parseInt(position_coded[1]),
    "position": parseInt(position_coded[2]),
    "sub_position": parseInt(position_coded[3]),
  };

  const response = await fetch('/api/floor/'+floor_id+'/io', {
    method: 'POST',
    headers: {
      'Content-Type': 'application/json'
    },
    body: JSON.stringify(data)
  });

  try {
    const result = await response.json();
    if (!response.ok) {
        console.warn(result.error);
    } else {
        factoryView_LoadIOs(window.G_FactoryID, floor_id);
        closedModal('io-new');
    }
  } catch(e) {
    console.error(e);
  }
}
