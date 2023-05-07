export const validateRequiredFields = (value: string) => {
  if (!value) {
    return "Required field";
  }
};

export const validateEmail = (value: string) => {
  if (!value) {
    return "Required";
  } else if (
    !/^(([^<>()[\]\\.,;:\s@"]+(\.[^<>()[\]\\.,;:\s@"]+)*)|(".+"))@((\[[0-9]{1,3}\.[0-9]{1,3}\.[0-9]{1,3}\.[0-9]{1,3}\])|(([a-zA-Z\-0-9]+\.)+[a-zA-Z]{2,}))$/i.test(
      value
    )
  ) {
    return "Email invalid adress";
  }
};
