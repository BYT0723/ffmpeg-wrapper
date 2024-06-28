# ffmpeg \
# 	-f x11grab -s 2240x1400 -r 30 -i $DISPLAY \
# 	-f v4l2 -video_size 320x240 -r 30 -i /dev/video0 \
# 	-f alsa -i default \
# 	-filter_complex "[0:v]scale=1920:1080[scaled];[scaled][1:v]overlay=W-w-100:H-h" \
# 	-preset ultrafast -tune zerolatency -threads $(echo "$(nproc)/2" | bc) output.mkv

desktop=$DISPLAY
desktop_size=2240x1400
desktop_rate=144
# echo $(xrandr)
camera_device=/dev/video0
camera_size=320x240
camera_rate=30
camera_position_x=W-w-100
camera_position_y=H-h
# echo $(v4l2-ctl --list-formats-ext -d "$CAMERA")

ffmpeg \
	-f x11grab -s 2240x1400 -r 144 -i $DISPLAY \
	-f v4l2 -video_size 320x240 -r 30 -i /dev/video0 \
	-f pulse -i "$(pactl get-default-sink).monitor" \
	-f pulse -i "$(pactl get-default-source)" \
	-filter_complex "[0:v]scale=1920:1200[scaled];[scaled][1:v]overlay=W-w-100:H-h" \
	-c:v libx264 -c:a aac -preset superfast -tune zerolatency -threads "$(echo "$(nproc)/2" | bc)" -bufsize 2M -rtbufsize 100M -f mp4 output.mp4

audio_output_device="$(pactl get-default-sink).monitor"
audio_input_device="$(pactl get-default-source)"

threads=$(echo "$(nproc)/2" | bc)

ffmpeg \
	-f x11grab -s $desktop_size -r $desktop_rate -i $desktop \
	-f v4l2 -video_size $camera_size -r $camera_rate -i "$camera_device" \
	-f pulse -i "$audio_output_device" \
	-f pulse -i "$audio_output_device" \
	-filter_complex "[0:v]scale=1920:1200[scaled];[scaled][1:v]overlay=$camera_position_x:$camera_position_y" \
	-c:v libx264 -c:a aac -preset superfast -tune zerolatency -threads $threads -bufsize 2M -rtbufsize 100M -f mp4 output.mp4
