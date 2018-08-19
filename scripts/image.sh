#!/bin/bash

src_dir="/images/post/$(basename $(pwd))"

for i in *.jpg; do 
  if [[ ${i} = *"_small.jpg" ]]; then
    continue;
  fi
  target="$(basename $i .jpg)_small.jpg"
  echo "{{<figure src=\"${src_dir}/${target}\" link=\"${src_dir}/${i}\">}}";
  if [[ -r $target  ]] ; then
    continue;
  fi
  convert $i -resize 582 ${target};
done

for i in *.mp4; do
  target="$(basename $i .mp4).webm"
  cat << EOF
<video autoplay loop muted playsinline>
  <source src="/${src_dir}/${target}" type="video/webm" />
</video>
EOF
  if [[ -r $target  ]] ; then
    continue
  fi
  avconv -t 6 -i ${i} -vf scale=582:-1 -c:v libvpx-vp9 -b:v 0 -crf 41 -an $target;
done
