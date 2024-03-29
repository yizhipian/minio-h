// This file is part of MinIO Console Server
// Copyright (c) 2022 MinIO, Inc.
//
// This program is free software: you can redistribute it and/or modify
// it under the terms of the GNU Affero General Public License as published by
// the Free Software Foundation, either version 3 of the License, or
// (at your option) any later version.
//
// This program is distributed in the hope that it will be useful,
// but WITHOUT ANY WARRANTY; without even the implied warranty of
// MERCHANTABILITY or FITNESS FOR A PARTICULAR PURPOSE.  See the
// GNU Affero General Public License for more details.
//
// You should have received a copy of the GNU Affero General Public License
// along with this program.  If not, see <http://www.gnu.org/licenses/>.

import React from "react";
import { InputBox } from "mds";
import { setSecretKey } from "./AddUsersSlice";
import { useSelector } from "react-redux";
import { AppState, useAppDispatch } from "../../../store";

const PasswordSelector = () => {
  const dispatch = useAppDispatch();
  const secretKey = useSelector(
    (state: AppState) => state.createUser.secretKey,
  );

  return (
    <InputBox
      id="standard-multiline-static"
      name="standard-multiline-static"
      type="password"
      label="Password"
      value={secretKey}
      onChange={(e: React.ChangeEvent<HTMLInputElement>) => {
        dispatch(setSecretKey(e.target.value));
      }}
      autoComplete="current-password"
    />
  );
};
export default PasswordSelector;
