/*
 * Copyright 2019 Google LLC
 *
 * Licensed under the Apache License, Version 2.0 (the "License");
 * you may not use this file except in compliance with the License.
 * You may obtain a copy of the License at
 *
 *      http://www.apache.org/licenses/LICENSE-2.0
 *
 * Unless required by applicable law or agreed to in writing, software
 * distributed under the License is distributed on an "AS IS" BASIS,
 * WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
 * See the License for the specific language governing permissions and
 * limitations under the License.
 */

/**
 * changeSelectAny handles checkbox #select-anything state change.
 * disables all autoselected checkbox when checked #select-anything
 */
function changeSelectAnything() {
  let e = document.getElementById("select-anything");
  let cbs = document.querySelectorAll(".autoselected");
  for (let i = 0; i < cbs.length; ++i) {
    cbs[i].disabled = e.checked;
  }
}

/**
 * init ...
 */
function init() {
  document.getElementById("select-anything").onchange = changeSelectAnything;
  document.getElementById("disagree").onclick = () => {
    document.getElementById("reject-form").submit();
  };
}

window.onload = init;
