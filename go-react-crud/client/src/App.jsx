
function App() {

  return (
    <>
    <h1>Hello world React!</h1>
     <button onClick={ async ()=>
     {
       const response = await fetch('/users')
       const data = await response.json()
       console.log(data)
     }}> Obtener datos </button>
    </>
   
  )
}

export default App
