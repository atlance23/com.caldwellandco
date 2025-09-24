function Header() {

    // Header Component
    return (
        <>
            <div className="headerContainer">
                <div className="headerLeft">
                    <img width={125} src="logo.avif" alt="Caldwell & Co Logo" id="logo" />
                </div>
                <div className="headerRight">
                    <a href="#" className="menuItem">Home</a>
                    <a href="#" className="menuItem">Blog</a>
                    <a href="#" className="menuItem">Watch Live</a>
                    <a href="#" className="menuItem">Contact</a>
                    <a href="#" className="menuItem">Shop</a>
                </div>
            </div>
        </>
    )
}

export default Header;