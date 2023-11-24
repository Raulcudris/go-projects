
function App() {

  const handleSubmit =() =>{

  }

  return (
    <div>
      <form onSubmit={()=> handleSubmit()}>
        <input type="name" placeholder="Coloca tu nombre"/>
        <button>Guardar</button>
      </form>
    </div>   
  )
}

export default App
