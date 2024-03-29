import React, {Component} from 'react'

const Hoc = (OriginalComponent) => {

    //here u can see that we can add common functioanlity/state etc to components through HOC
    //here any component A & B that exports using HOC has a prop name due to the HOC
    
    //u need to do {...this.props} because in App.js when u call the <A age={22}> since its exported as Hoc(A)
    // the prop age is not a member of A but gets passed to the HOC .. u need to pass it down to the A via {...this.props}
    class NewComponent extends Component {
      render() {
        return (
          <OriginalComponent name="saishwar" {...this.props}/>
        )
      }
    }

    return NewComponent
    
}

export default Hoc