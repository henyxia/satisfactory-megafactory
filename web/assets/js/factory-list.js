async function modalFactoryNew_Send() {
    data = {
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
