import React from 'react';
import { Link } from 'react-router-dom';
import Yene from '../assets/Yene.png';
import Sean from '../assets/Sean.png';
import Aramis from '../assets/Aramis.png';
import ProductPic from '../assets/ProductPic.png'; // Ensure the logo path is correct

import '../styles/home.css'; // Ensure you have this CSS file

function Home() {
    return (
        <body className="font-kanit text-lg text-white bg-gray-800">
            <main>
                <section class="relative overflow-hidden md:py-48 py-40 bg-Purple/5" id="home">
                    <div class="container relative mt-8">
                        <div class="flex space-x-4 justify-between">
                            <div class="">
                                <h1 class="font-semibold lg:leading-normal leading-normal tracking-wide text-4xl lg:text-5xl mb-5">Use One Modern, Fast, and Easy to Use Tool</h1>
                                <p class="text-slate-400 text-lg max-w-xl"> Bravus has the ability to oragize appointments, make spreadsheets, make tables and graphics, and all in a simple to use dashboard</p>
                                <div class="mt-6">
                                    <Link
                                    to="/signup" 
                                    className="
                                    relative flex items-center justify-center h-[50px] sm:w-[50px] md:w-[100px] lg:w-[150px] xl:w-[200px] 2xl:w-[200px] text-lg group right-8 rounded-3xl overflow-visible transition-all duration-500
                                    "
                                    >
                                    <span className="relative w-[110px] z-10 block px-5 py-3 overflow-hidden font-medium leading-tight text-white transition-colors duration-300 ease-out border-1 border-black rounded-lg group-hover:text-white">
                                        <span className="absolute inset-0 w-full h-full px-5 py-3 rounded-lg bg-gray-800"></span>
                                        <span className="absolute left-0 w-48 h-48 -ml-2 transition-all duration-300 origin-top-right -rotate-90 -translate-x-full translate-y-12 bg-Purple group-hover:-rotate-180 ease"></span>
                                        <span className="relative">Register</span>
                                    </span>
                                    <span className="absolute bottom-[2px] right-10 w-[110px] h-10 -mb-1 -mr-1 transition-all duration-200 ease-linear bg-Purple rounded-lg group-hover:mb-1 group-hover:mr-1"></span>
                                    </Link>
                                </div>
                            </div>
                            <div class="lg:ms-8">
                                <div class="relative">
                                    <img src={ProductPic} class="relative scale-120 border-4 rounded-xl border-Purple shadow-2xl object-fit h-[450px] w-[2000px]  m-[5px]" alt=""/>
                                    <div class="overflow-hidden absolute md:size-[500px] size-[400px] bg-gradient-to-tl to-teal-500/20 via-teal-500/70 from-teal-500 bottom-1/2 translate-y-1/2 md:start-0 start-1/2 ltr:md:translate-x-0 ltr:-translate-x-1/2 rtl:md:translate-x-0 rtl:translate-x-1/2 -z-1 shadow-md shadow-teal-500/10 rounded-full"></div>
                                </div>
                            </div>
                        </div>
                    </div>
                </section>

                <section id="about" className="relative md:py-24 py-16 bg-gray-700 text-center rounded-full">
                    <div className="container mx-auto">
                        <div className="grid md:grid-cols-1 grid-cols-1 items-center gap-[30px]">
                            <div>
                                <h2 className="lg:text-[48px] font-bold mb-4 text-purple-500">Why Bravus?</h2>
                                <p className="text-gray-400 max-w-xl mx-auto lg:text-[20px]">
                                    We aimed to create an app specifically for service-based businesses, such as salons and barbershops, to streamline their operations seamlessly. Our app integrates with the Google Sheets API to manage data efficiently and features an inbuilt appointment system that allows users to schedule and track appointments effortlessly. The app also includes advanced graphics and tools to help business owners manage their schedules and track their performance intuitively.
                                </p>
                            </div>
                        </div>
                    </div>
                </section>

                <section id="Team" className="relative md:py-24 py-16 text-center">
                    <div className="container mx-auto">
                        <div className="grid grid-cols-1">
                            <h2 className="lg:text-[48px] font-bold mb-4 text-purple-500">The Team</h2>
                            <p className="text-gray-400 max-w-xl mx-auto mb-8 lg:text-[20px]">
                                Meet the talented team behind Bravus. Together, we have brought our final project for our Software Engineering certificate to life, combining our skills and expertise to create an exceptional appointment scheduling experience.
                            </p>
                            <div className="grid md:grid-cols-3 gap-8">
                                <div className="team-member">
                                    <div className="team-member-photo">
                                        <Link to="https://www.linkedin.com/in/sean-cardona-bb1594166?trk=contact-info">
                                            <img src={Sean} alt="Sean A. Cardona" />
                                        </Link>
                                        <div className="peel-effect"></div>
                                        <p className="team-member-name">Sean A. Cardona</p>
                                    </div>
                                </div>
                                <div className="team-member">
                                    <div className="team-member-photo">
                                        <Link to="https://www.linkedin.com/in/aramis-martinez-a1a507296/">
                                            <img src={Aramis} alt="Aramis Martinez" />
                                        </Link>
                                        <div className="peel-effect"></div>
                                        <p className="team-member-name">Aramis Martinez</p>
                                    </div>
                                </div>
                                <div className="team-member">
                                    <div className="team-member-photo">
                                        <Link to="https://www.linkedin.com/in/yeneishla-santiago-958b63254/">
                                            <img src={Yene} alt="Yeneishla Santiago" />
                                        </Link>
                                        <div className="peel-effect"></div>
                                        <p className="team-member-name">Yeneishla Santiago</p>
                                    </div>
                                </div>
                            </div>
                        </div>
                    </div>
                </section>

            </main>
        </body>
    );
}

export default Home;