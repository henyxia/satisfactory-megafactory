async function modalFactoryDelete_Send(id) {
    const response = await fetch('/api/factory/'+id, {
        method: 'DELETE',
        headers: {
          'Content-Type': 'application/json'
        },
      });

    try {
        const result = await response.json();
        console.log(result);
        window.location.pathname = '/';
    } catch(e) {
        console.warn(e);
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
      console.log(result);
      window.location.reload();
  } catch(e) {
      console.warn(e);
  }
}
