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
