import React from 'react'

const E_ComponentProps = () => {
  return (
    <div>
        <Button/>
        <Button2 type='submit' autoFocus={true} defaultValue="test" xyz="lll"/>
    </div>
  )
}


//this helps to accept all the native button properties as props
type ButtonProps = React.ComponentPropsWithoutRef<'button'>

// if u want to have some aditional props
type ButtonProps2 = React.ComponentPropsWithoutRef<'button'> &{
    name: string,
    age: number
} 

// there is also (check their usecases)
type ButtonProps3 = React.ComponentPropsWithRef<'button'>
type ButtonProps4 = React.ComponentProps<'button'>


const Button = ({ type, autoFocus, onClick}: ButtonProps) => {
    return (
        <>
            <button></button>
        </>
    )
}



// **************** 2. Spread/rest operator ******************8


type Button2Props = React.ComponentPropsWithoutRef<'button'>
& {
    xyz: string
}

const Button2 = ({ type, ...rest }: Button2Props) => {
    console.log(rest) // {autoFocus: true, defaultValue: 'test', xyz: 'lll'}
    return (
        <>
            <button type={type} {...rest}> Button 2</button>
        </>
    )
}

export default E_ComponentProps