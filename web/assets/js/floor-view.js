async function modalFloorDelete_Send(factory_id, floor_id, redirect_url) {
    const response = await fetch('/api/factory/'+factory_id+'/floor/'+floor_id, {
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
            window.location.pathname = redirect_url;
        }
    } catch(e) {
        console.error(e);
    }
}

async function modalFloorEdit_Send(factory_id, floor_id) {
    data = {
        "id": floor_id,
        "name": getInputValue("modal-floor-new-name"),
        "level": getInputValueInt("modal-floor-new-level"),
        "status": getInputValue("modal-floor-new-status"),
        "factory_id": factory_id,
    };

    const response = await fetch('/api/factory/'+factory_id+'/floor/'+floor_id, {
        method: 'PUT',
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
