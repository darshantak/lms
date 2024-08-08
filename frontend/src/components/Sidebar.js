import React, { Fragment, useState } from "react";
import { IoIosArrowForward, IoIosArrowDown } from "react-icons/io";

const Sidebar = ({open,setOpen}) => {
//   const [open, setOpen] = useState(true);
  const [activeSubmenu, setActiveSubmenu] = useState(null);
  const Menus = [
    { title: "User Management", route: "users"},
    { title: "Trainings", route: "trainings" },
    { title: "Content", route: "content" },
    {title: "Assignments", route: "assignments"}
  ];
  const toggleSubMenu = (index) => {
    setActiveSubmenu(activeSubmenu === index ? null : index);
  };
  return (
    <div className="flex fixed">
      <div
        className={`${
          open ? "w-72" : "w-10"
        } bg-yellow-700 h-screen p-5 pt-8 relative duration-300`}
      >
        <img
          src={`${process.env.PUBLIC_URL + "/control.png"}`}
          className={`absolute cursor-pointer -right-3 top-9 w-7 border-dark-purple
           border-2 rounded-full  ${!open && "rotate-180"}`}
          onClick={() => setOpen(!open)}
        />
        <ul className="pt-10">
          {Menus.map((Menu, index) => (
            <div key={index}>
              <li
                key={index}
                className={`flex rounded-md p-2 cursor-pointer hover:bg-light-white text-gray-300 text-lg items-center gap-x-4 ${
                  Menu.gap ? "mt-9" : "mt-2"
                } ${index === 0 && "bg-light-white"}`}
                onClick={() => Menu.submenu && toggleSubMenu(index)}
              >
                <span
                  className={`${!open && "hidden"} origin-left duration-200`}
                >
                  {Menu.title}
                </span>
                {Menu.submenu && (
                  <span
                    className={`${
                      !open && "hidden"
                    } ml-auto origin-left duration-200`}
                  >
                    {activeSubmenu === index ? (
                      <IoIosArrowDown />
                    ) : (
                      <IoIosArrowForward />
                    )}
                  </span>
                )}

                {/* TODO: For icon another span will come here */}
                {/* {<span className={`${!open && "hidden"} origin-left duration-200`}>
                {Menu.title}
              </span>} */}
              </li>
              {Menu.submenu && (
                <ul
                  className={`ml-4 ${
                    activeSubmenu === index ? "block" : "hidden"
                  } ease-in-out duration-300`}
                >
                  {Menu.submenu.map((subItem, subIndex) => (
                    <li
                      key={subIndex}
                      className="text-gray-200 text-sm p-2 hover:bg-gray-700 rounded-md"
                    >
                      <a href="#">{subItem}</a>
                    </li>
                  ))}
                </ul>
              )}
            </div>
          ))}
        </ul>
      </div>
    </div>
  );
};

export default Sidebar;
